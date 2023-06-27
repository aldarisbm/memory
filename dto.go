package memory

import (
	"encoding/json"
	"errors"
	"github.com/aldarisbm/memory/datasource"
	"github.com/aldarisbm/memory/datasource/boltdb"
	"github.com/aldarisbm/memory/datasource/sqlite"
	"github.com/aldarisbm/memory/embeddings"
	"github.com/aldarisbm/memory/embeddings/local"
	oai "github.com/aldarisbm/memory/embeddings/openai"
	"github.com/aldarisbm/memory/vectorstore"
	"github.com/aldarisbm/memory/vectorstore/heisenberg"
	pc "github.com/aldarisbm/memory/vectorstore/pinecone"
)

var (
	ErrInvalidVectorStoreType = errors.New("invalid vector store type")
	ErrInvalidEmbedderType    = errors.New("invalid embedder type")
	ErrInvalidDataSourceType  = errors.New("invalid data source type")
)

type DTO struct {
	VS      vectorstore.Converter `json:"vs"`
	VSType  string                `json:"vsType"`
	Emb     embeddings.Converter  `json:"emb"`
	EmbType string                `json:"embType"`
	DS      datasource.Converter  `json:"ds"`
	DSType  string                `json:"dsType"`
}

func (d *DTO) MarshallJSON() ([]byte, error) {
	return json.Marshal(d)
}

func (d *DTO) UnmarshalJSON(b []byte) error {
	// Unmarshal to a temporary struct with same structure, but where MyField is a json.RawMessage
	tmpStruct := struct {
		VS      json.RawMessage
		VsType  string
		Emb     json.RawMessage
		EmbType string
		DS      json.RawMessage
		DSType  string
	}{}

	if err := json.Unmarshal(b, &tmpStruct); err != nil {
		return err
	}

	d.VSType = tmpStruct.VsType
	d.EmbType = tmpStruct.EmbType
	d.DSType = tmpStruct.DSType

	switch tmpStruct.VsType {
	case "heisenberg":
		d.VS = &heisenberg.DTO{}
	case "pinecone":
		d.VS = &pc.DTO{}
	default:
		return ErrInvalidVectorStoreType
	}
	if err := json.Unmarshal(tmpStruct.VS, d.VS); err != nil {
		return err
	}

	switch tmpStruct.EmbType {
	case "openai":
		d.Emb = &oai.DTO{}
	case "local":
		d.Emb = &local.DTO{}
	default:
		return ErrInvalidEmbedderType
	}
	if err := json.Unmarshal(tmpStruct.Emb, d.Emb); err != nil {
		return err
	}

	switch tmpStruct.DSType {
	case "boltdb":
		d.DS = &boltdb.DTO{}
	case "sqlite":
		d.DS = &sqlite.DTO{}
	default:
		return ErrInvalidDataSourceType
	}

	if err := json.Unmarshal(tmpStruct.DS, d.DS); err != nil {
		return err
	}

	return nil
}
