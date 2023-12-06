package db

import (
	"encoding/json"

	"github.com/tuckersn/chatbackend/openai"
)

type EmbeddingsRecord struct {
	Id         int32     `db:"id"`
	Embeddings []float64 `db:"embeddings"`
	TableName  string    `db:"table_name"`
	TableCol   string    `db:"table_col"`
	TableId    int32     `db:"table_id"`
	Model      string    `db:"model"`
	Created    int64     `db:"created"`
}

/*
OpenAI Embeddings
only enabled if pgvector is enabled and OpenAI key is provided
https://platform.openai.com/docs/guides/embeddings

table name is plural since every record contains 1536 floats
*/
func TableInitOpenAIEmbeddings(context TableInitContext) {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS openai_embeddings (
			id SERIAL PRIMARY KEY,
			embeddings VECTOR(1536) NOT NULL,
			table_name TEXT NOT NULL,
			table_col TEXT NOT NULL,
			table_id INTEGER NOT NULL,
			model TEXT NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			UNIQUE(table_name, table_col, table_id)
		);
	`)

	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_openai_embeddings_embeddings ON openai_embeddings (embeddings);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_openai_embeddings_table_name ON openai_embeddings (table_name);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_openai_embeddings_table_col ON openai_embeddings (table_col);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_openai_embeddings_table_id ON openai_embeddings (table_id);`)
}

func GetOrVectorizeRecordCustom(
	tableName string,
	tableCol string,
	tableId int32,
	idFieldName string,
	model string,
) (EmbeddingsRecord, error) {

	var embeddingsRecord EmbeddingsRecord
	err := Con.QueryRow(`
		SELECT * FROM openai_embeddings
		WHERE table_name = $1 AND table_col = $2 AND table_id = $3
	`, tableName, tableCol, tableId).Scan(&embeddingsRecord)
	if err != nil {
		panic(err)
	}

	if &embeddingsRecord != nil {
		return embeddingsRecord, nil
	}

	columnContent, err := Con.Exec(`
		SELECT $1 FROM $2 WHERE $4 = $3
	`, tableCol, tableName, tableId, idFieldName)
	if err != nil {
		panic(err)
	}

	jsonValue, err := json.Marshal(columnContent)
	if err != nil {
		panic(err)
	}

	embeddings, err := openai.GenerateEmbeddings(string(jsonValue))
	if err != nil {
		panic(err)
	}

	_, err = Con.Exec(`
		INSERT INTO openai_embeddings (embeddings, table_name, table_col, table_id)
		VALUES ($1, $2, $3, $4)
	`, embeddings.Vectors, tableName, tableCol, tableId)
	if err != nil {
		panic(err)
	}

	return EmbeddingsRecord{
		Embeddings: embeddings.Vectors,
		TableName:  tableName,
		TableCol:   tableCol,
		TableId:    tableId,
		Model:      model,
	}, nil

}

func GetOrVectorizeRecord(
	tableName string,
	tableCol string,
	tableId int32,
) (EmbeddingsRecord, error) {
	return GetOrVectorizeRecordCustom(tableName, tableCol, tableId, "id", "text-embedding-ada-002")
}
