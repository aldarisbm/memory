package kvstore

import (
	"github.com/aldarisbm/memory"
	"github.com/google/uuid"
)

type Store interface {
	Get(key uuid.UUID) (memory.Memory, error)
	Set(key uuid.UUID, value memory.Memory) error
}
