package boltdb

import (
	"encoding/json"
	"fmt"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
)

type LocalStore struct {
	db         *bolt.DB
	bucketName string
}

func NewLocalStore(opts ...CallOptions) *LocalStore {
	o := applyCallOptions(opts, options{
		path:   "localdb",
		bucket: "ltm",
	})
	dbm, err := bolt.Open(o.path, 0600, nil)
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

	ls := &LocalStore{
		db:         dbm,
		bucketName: o.bucket,
	}
	return ls
}

func (l *LocalStore) GetDocument(id uuid.UUID) (*shared.Document, error) {
	var doc shared.Document
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
		return nil, fmt.Errorf("getting document: %s", err)
	}
	return &doc, nil
}

func (l *LocalStore) GetDocuments(ids []uuid.UUID) ([]*shared.Document, error) {
	var docs []*shared.Document
	for _, id := range ids {
		doc, err := l.GetDocument(id)
		if err != nil {
			return nil, fmt.Errorf("getting document: %s", err)
		}
		docs = append(docs, doc)
	}
	return docs, nil
}

func (l *LocalStore) StoreDocument(document *shared.Document) error {
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
