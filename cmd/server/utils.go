package main

import (
	"github.com/aldarisbm/memory/internal"
	bolt "go.etcd.io/bbolt"
)

func createKVStore() *bolt.DB {
	path := fmt.Sprintf("%s/kvstore", internal.CreateMemoryFolderInHomeDir())
	mode := 0600
	bucket := "service"
	dbm, err := bolt.Open(path, os.FileMode(mode), nil)
	if err != nil {
		panic(err)
	}
	err = dbm.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(bucket))
		if err != nil {
			panic(err)
		}
		return nil
	})
	return dbm
}
