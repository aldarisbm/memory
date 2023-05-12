package main

import (
	"github.com/aldarisbm/ltm/pkg/datasource/boltdb"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
	"time"
)

func main() {

	id := uuid.New()
	doc := shared.Document{
		ID:         id,
		Text:       "mi mama no me mima",
		CreatedAt:  time.Now(),
		LastReadAt: time.Now(),
	}
	localStore := boltdb.NewLocalStore()

	err := localStore.StoreDocument(&doc)
	if err != nil {
		panic(err)
	}
	s, err := localStore.GetDocument(id)
	if err != nil {
		panic(err)
	}
	println(s.Text)
}
