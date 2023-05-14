package main

import (
	"github.com/aldarisbm/ltm/pkg/datasource/sqlite"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/google/uuid"
	"time"
)

func main() {

	id := uuid.New()
	doc := shared.Document{
		ID:         id,
		Text:       "mi mama si me mima",
		CreatedAt:  time.Now(),
		LastReadAt: time.Now(),
	}
	localStore := sqlite.NewLocalStorer()
	defer localStore.Close()

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
