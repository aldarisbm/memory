package sqliteds

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/user"
	"time"

	"github.com/aldarisbm/memory/datasource"
	"github.com/aldarisbm/memory/internal"
	"github.com/aldarisbm/memory/types"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

type sqliteds struct {
	db   *sql.DB
	path string
	DTO  *DTO
}

// New returns a new local storer
// if path is empty, it will default to $HOME/memory/memory.db
func New(opts ...CallOptions) *sqliteds {
	o := applyCallOptions(opts)
	if o.path == "" {
		o.path = fmt.Sprintf("%s/%s", internal.CreateMemoryFolderInHomeDir(), fmt.Sprintf("%s.db", internal.Generate(10)))
	}
	db, err := createTable(o.path)
	if err != nil {
		log.Fatal(err)
	}

	ls := &sqliteds{
		db:   db,
		path: o.path,
		DTO: &DTO{
			Path: o.path,
		},
	}
	return ls
}

// Close closes the local storer
func (s *sqliteds) Close() error {
	return s.db.Close()
}

// GetDocument returns the document with the given id
func (s *sqliteds) GetDocument(id uuid.UUID) (*types.Document, error) {
	var doc types.Document

	stmt, err := s.db.Prepare("SELECT id, user, text, created_at, last_read_at, vector, metadata FROM documents WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	var metadataBytes, vectorBytes []byte
	err = row.Scan(&doc.ID, &doc.User, &doc.Text, &doc.CreatedAt, &doc.LastReadAt, &vectorBytes, &metadataBytes)
	if err != nil {
		return nil, fmt.Errorf("scanning document: %s", err)
	}
	if err := json.Unmarshal(metadataBytes, &doc.Metadata); err != nil {
		return nil, fmt.Errorf("unmarshaling metadata: %s", err)
	}
	if err := json.Unmarshal(vectorBytes, &doc.Vector); err != nil {
		return nil, fmt.Errorf("unmarshaling vector: %s", err)
	}
	err = l.UpdateLastReadAt(&doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// GetDocuments returns the documents with the given ids
func (s *sqliteds) GetDocuments(ids []uuid.UUID) ([]*types.Document, error) {
	// TODO should probably do this in a single query
	var docs []*types.Document

	for _, id := range ids {
		doc, err := s.GetDocument(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// StoreDocument stores the given document
// We marshal the vector and metadata to []byte and store them as blobs
func (s *sqliteds) StoreDocument(doc *types.Document) error {
	stmt, err := s.db.Prepare("INSERT INTO documents (id, user,  text, created_at, last_read_at, vector, metadata) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("preparing statement: %s", err)
	}

	vector, err := json.Marshal(doc.Vector)
	if err != nil {
		return fmt.Errorf("marshaling vector: %s", err)
	}
	metadata, err := json.Marshal(doc.Metadata)
	if err != nil {
		return fmt.Errorf("marshaling metadata: %s", err)
	}
	res, err := stmt.Exec(doc.ID, doc.User, doc.Text, doc.CreatedAt, doc.LastReadAt, vector, metadata)
	if err != nil {
		return fmt.Errorf("inserting document: %s", err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert id: %s", err)
	}
	return nil
}


// UpdateLastReadAt updates the last read at timestamp for the given document
func (l *localStorer) UpdateLastReadAt(doc *types.Document) error {
	if time.Now().After(doc.LastReadAt) {
		doc.LastReadAt = time.Now() // Update prevTime with the new current time
	} else {
		return fmt.Errorf("error in updating LastReadAt time")
	}
	return nil
}

// Ensure localStorer implements DataSourcer
var _ datasource.DataSourcer = (*localStorer)(nil)
// GetDTO returns the DTO of the local storer
func (s *sqliteds) GetDTO() datasource.Converter {
	return s.DTO
}

// Ensure sqliteds implements DataSourcer
var _ datasource.DataSourcer = (*sqliteds)(nil)

