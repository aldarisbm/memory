package main

import (
	"github.com/aldarisbm/ltm/pkg"
	"github.com/aldarisbm/ltm/pkg/shared"
	"github.com/golang/protobuf/ptypes/timestamp"
)

func main() {

	doc := shared.Document{
		ID:         "123321",
		Text:       "mi mama me mimaba y yo la mimaba",
		CreatedAt:  timestamp.Timestamp{},
		LastReadAt: timestamp.Timestamp{},
	}
	ltm := pkg.NewLTM()
	err := ltm.StoreDocument(&doc)
	if err != nil {
		panic(err)
	}

}
