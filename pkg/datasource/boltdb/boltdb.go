package boltdb

import (
	"encoding/json"
	"fmt"
	"github.com/aldarisbm/memory/pkg/datasource"
	"github.com/aldarisbm/memory/pkg/document"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type localStorer struct {
	db         *bolt.DB
	bucketName string
}

func NewLocalStorer(opts ...CallOptions) *localStorer {
	o := applyCallOptions(opts, options{
		path:   "localdb",
		bucket: "ltm",
		mode:   0600,
	})
	dbm, err := bolt.Open(o.path, o.mode, nil)
	if err != nil {
		panic(err)
	}
	err = dbm.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(o.bucket))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	ls := &localStorer{
		db:         dbm,
		bucketName: o.bucket,
	}
	return ls
}

func (l *localStorer) Close() error {
	return l.db.Close()
}

func (l *localStorer) GetDocument(id uuid.UUID) (*document.Document, error) {
	var doc document.Document
	err := l.db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(l.bucketName))
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
	return &doc, nil
}

func (l *localStorer) GetDocuments(ids []uuid.UUID) ([]*document.Document, error) {
	var docs []*document.Document
	for _, id := range ids {
		doc, err := l.GetDocument(id)
		if err != nil {
			return nil, fmt.Errorf("getting document: %s", err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func (l *localStorer) StoreDocument(document *document.Document) error {
	doc, err := json.Marshal(&document)
	if err != nil {
		return fmt.Errorf("marshaling document: %s", err)
	}
	err = l.db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(l.bucketName))
		err := b.Put([]byte(document.ID.String()), doc)
		return err
	})
	if err != nil {
		return fmt.Errorf("updating bolt db: %s", err)
	}
	return nil
}

// Ensure localStorer implements DataSourcer
var _ datasource.DataSourcer = (*localStorer)(nil)
