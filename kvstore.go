package memory

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/aldarisbm/memory/internal"
	bolt "go.etcd.io/bbolt"
)

type boltStore struct {
	db         *bolt.DB
	name       string
	bucketName string
}

func getStore() storer {
	path := fmt.Sprintf("%s/%s", internal.CreateMemoryFolderInHomeDir(), BoltDB)
	dbm, err := bolt.Open(path, 0600, nil)
	if err != nil {
		panic(err)
	}
	err = dbm.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(BucketName))
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		panic(err)
	}

	return &boltStore{
		db:         dbm,
		bucketName: BucketName,
		name:       BoltDB,
	}
}

func (b *boltStore) saveMemoryToStore(name string, mem *Memory) error {
	dto := &DTO{
		VS:  mem.vectorStore.GetDTO(),
		Emb: mem.embedder.GetDTO(),
		DS:  mem.datasource.GetDTO(),
	}

	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	if err := enc.Encode(dto); err != nil {
		return err
	}

	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))
		if v != nil {
			return fmt.Errorf("memory with name %s already exists", name)
		}

		return bucket.Put([]byte(name), buf.Bytes())
	})
}

func (b *boltStore) getMemoryFromStore(name string) (*Memory, error) {
	var dto *DTO

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))

		dec := gob.NewDecoder(bytes.NewReader(v))
		if err := dec.Decode(dto); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return &Memory{
		vectorStore: dto.VS.ToVectorStore(),
		embedder:    dto.Emb.ToEmbedder(),
		datasource:  dto.DS.ToDataSource(),
	}, nil
}

func (b *boltStore) close() error {
	return b.db.Close()
}

var _ storer = (*boltStore)(nil)
