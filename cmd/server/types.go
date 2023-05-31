package main

import "github.com/google/uuid"

type Memory struct {
	ID uuid.UUID `json:"id"`
}

type PostMemoryRes struct {
	ID uuid.UUID `json:"id"`
}
