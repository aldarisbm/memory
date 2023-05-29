package sqlitevss

import (
	"database/sql"
	"github.com/aldarisbm/memory/types"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/google/uuid"
	"github.com/mattn/go-sqlite3"
)

// sqliteVSS is a vector store that uses SQLite as the backend
type sqliteVSS struct {
	db *sql.DB
}

// NewSQLiteVSS returns a new sqliteVSS
func NewSQLiteVSS(options ...CallOptions) *sqliteVSS {
	sql.Register("sqlite3_with_extensions", &sqlite3.SQLiteDriver{
		Extensions: []string{
			"pkg/vectorstore/sqlitevss/plugins/vector0",
			"pkg/vectorstore/sqlitevss/plugins/vss0",
		},
	})
	db, err := createDatabase("test.db")
	if err != nil {
		panic(err)
	}

	return &sqliteVSS{
		db: db,
	}
}

// StoreVector stores the given Document
func (s *sqliteVSS) StoreVector(document *types.Document) error {

	return nil
}

// QuerySimilarity returns the k most similar documents to the given vector
func (s *sqliteVSS) QuerySimilarity(vector []float32, k int64) ([]uuid.UUID, error) {
	return nil, nil
}

// Close closes the sqliteVSS
func (s *sqliteVSS) Close() error {
	return nil
}

// Ensure that sqliteVSS implements VectorStorer
var _ vectorstore.VectorStorer = (*sqliteVSS)(nil)
