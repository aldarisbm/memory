package ltm

import "github.com/golang/protobuf/ptypes/timestamp"

type Document struct {
	Id         string              `json:"id"`
	Text       string              `json:"text"`
	CreatedAt  timestamp.Timestamp `json:"created_at"`
	LastReadAt timestamp.Timestamp `json:"last_read_at"`
}
