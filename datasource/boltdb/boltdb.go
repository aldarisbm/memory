package boltds

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"time"

	"github.com/aldarisbm/memory/datasource"
	"github.com/aldarisbm/memory/internal"
	"github.com/aldarisbm/memory/types"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"os"
)

// bucketName is always the same
const (
	bucketName = "ltm"
	mode       = os.ModePerm
)

type boltds struct {
	path string
	db   *bolt.DB
	DTO  *DTO
}

// New returns a new local storer
// if path is empty, it will default to $HOME/memory/boltdb
func New(opts ...CallOptions) *boltds {
	o := applyCallOptions(opts)
	if o.path == "" {
		o.path = fmt.Sprintf("%s/%s", internal.CreateMemoryFolderInHomeDir(), internal.Generate(10))
	}
	dbm, err := bolt.Open(o.path, mode, nil)
	if err != nil {
		panic(err)
	}
	err = dbm.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	ls := &boltds{
		db:   dbm,
		path: o.path,
		DTO: &DTO{
			Path: o.path,
		},
	}
	return ls
}

// Close closes the local storer
func (b *boltds) Close() error {
	return b.db.Close()
}

// GetDocument returns the document with the given id
func (b *boltds) GetDocument(id uuid.UUID) (*types.Document, error) {
	var doc types.Document
	err := b.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		v := b.Get([]byte(id.String()))
		err := json.Unmarshal(v, &doc)
		if err != nil {
			return fmt.Errorf("unmarshaling document: %s", err)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	err = l.UpdateLastReadAt(&doc)
	if err != nil {
		return nil, err
	}
	return &doc, nil
}

// GetDocuments returns the documents with the given ids
func (b *boltds) GetDocuments(ids []uuid.UUID) ([]*types.Document, error) {
	var docs []*types.Document
	for _, id := range ids {
		doc, err := b.GetDocument(id)
		if err != nil {
			return nil, fmt.Errorf("getting document: %s", err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

// StoreDocument stores the given document
// We use a k/v store key being uuid and value being []byte of Document
func (b *boltds) StoreDocument(document *types.Document) error {
	doc, err := json.Marshal(&document)
	if err != nil {
		return fmt.Errorf("marshaling document: %s", err)
	}
	err = b.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(bucketName))
		err := b.Put([]byte(document.ID.String()), doc)
		return err
	})
	if err != nil {
		return fmt.Errorf("updating bolt db: %s", err)
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
func (b *boltds) GetDTO() datasource.Converter {
	return b.DTO
}

// Ensure boltds implements DataSourcer
var _ datasource.DataSourcer = (*boltds)(nil)

