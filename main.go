package main

import (
	"github.com/aldarisbm/ltm/pkg/datasource/boltdb"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {

	doc := shared.Document{
		ID:         "321",
		Text:       "mi mama no me mima",
		CreatedAt:  timestamp.Timestamp{},
		LastReadAt: timestamp.Timestamp{},
	}

	localStore := boltdb.NewLocalStore(
		boltdb.WithPath("boltdb"),
		boltdb.WithBucket("ltm"),
	)

	err := localStore.StoreDocument(&doc)
	if err != nil {
		panic(err)
	}
	s, err := localStore.GetDocument("321")
	if err != nil {
		panic(err)
	}
	println(s.Text)
}
