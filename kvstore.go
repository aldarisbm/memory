package memory

import (
	"fmt"
	"github.com/aldarisbm/memory/internal"
	"github.com/aldarisbm/memory/types"
	"github.com/google/uuid"
	bolt "go.etcd.io/bbolt"
	"log"
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

	vsDTO := mem.vectorStore.GetDTO()
	embDTO := mem.embedder.GetDTO()
	dsDTO := mem.datasource.GetDTO()
	dto := DTO{
		VS:      vsDTO,
		VSType:  vsDTO.GetType(),
		Emb:     embDTO,
		EmbType: embDTO.GetType(),
		DS:      dsDTO,
		DSType:  dsDTO.GetType(),
	}

	bs, err := dto.MarshallJSON()
	if err != nil {
		return err
	}

	return b.db.Update(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))
		if v != nil {
			return fmt.Errorf("memory with name %s already exists", name)
		}

		return bucket.Put([]byte(name), bs)
	})
}

func (b *boltStore) getMemoryFromStore(name string) (*Memory, error) {
	var dto DTO

	err := b.db.View(func(tx *bolt.Tx) error {
		bucket := tx.Bucket([]byte(BucketName))
		v := bucket.Get([]byte(name))

		if err := dto.UnmarshalJSON(v); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		log.Default().Println(err)
		return nil, err
	}
	return &Memory{
		vectorStore: dto.VS.ToVectorStore(),
		embedder:    dto.Emb.ToEmbedder(),
		datasource:  dto.DS.ToDataSource(),
		cache:       make(map[uuid.UUID]*types.Document),
	}, nil
}

func (b *boltStore) close() error {
	return b.db.Close()
}

var _ storer = (*boltStore)(nil)
