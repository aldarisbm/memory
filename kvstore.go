package memory

import (
	"encoding/json"
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
	m, err := json.Marshal(mem)
	if err != nil {
		return err
	}

	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))
		if v != nil {
			return fmt.Errorf("memory with name %s already exists", name)
		}

		return bucket.Put([]byte(name), m)
	})
}

func (b *boltStore) getMemoryFromStore(name string) (*Memory, error) {
	var mem *Memory

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))

		if err := json.Unmarshal(v, &mem); err != nil {
			return err
		}
		if mem == nil {
			return fmt.Errorf("memory with name %s does not exist", name)
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return mem, nil
}

func (b *boltStore) close() error {
	return b.db.Close()
}

var _ storer = (*boltStore)(nil)
