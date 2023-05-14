package sqlite

import (
	"database/sql"
	"fmt"
	"github.com/aldarisbm/ltm/pkg/datasource"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
	"log"
)

type LocalStorer struct {
	db   *sql.DB
	path string
}

func NewLocalStorer(opts ...CallOptions) *LocalStorer {
	o := applyCallOptions(opts, options{
		path: "localdb.db",
	})
	db, err := createTable(o.path)
	if err != nil {
		log.Fatal(err)
	}

	ls := &LocalStorer{
		db:   db,
		path: o.path,
	}
	return ls
}

func (l *LocalStorer) Close() error {
	return l.db.Close()
}

func (l *LocalStorer) GetDocument(id uuid.UUID) (*shared.Document, error) {
	var doc shared.Document

	di := id.String()
	fmt.Println("id", di)
	stmt, err := l.db.Prepare("SELECT id, text, created_at, last_read_at FROM documents WHERE id=?")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(id)
	err = row.Scan(&doc.ID, &doc.Text, &doc.CreatedAt, &doc.LastReadAt)
	if err != nil {
		return nil, fmt.Errorf("scanning document: %s", err)
	}
	return &doc, nil
}

func (l *LocalStorer) GetDocuments(ids []uuid.UUID) ([]*shared.Document, error) {
	// TODO should probably do this in a single query
	var docs []*shared.Document

	for _, id := range ids {
		doc, err := l.GetDocument(id)
		if err != nil {
			return nil, err
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func (l *LocalStorer) StoreDocument(doc *shared.Document) error {
	stmt, err := l.db.Prepare("INSERT INTO documents (id, text, created_at, last_read_at) VALUES (?, ?, ?, ?)")
	if err != nil {
		return fmt.Errorf("preparing statement: %s", err)
	}

	res, err := stmt.Exec(doc.ID, doc.Text, doc.CreatedAt, doc.LastReadAt)
	if err != nil {
		return fmt.Errorf("inserting document: %s", err)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return fmt.Errorf("getting last insert id: %s", err)
	}
	return nil
}

// Ensure LocalStorer implements DataSourcer
var _ datasource.DataSourcer = (*LocalStorer)(nil)
