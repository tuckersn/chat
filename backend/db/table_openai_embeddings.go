package db

func DBInitializeOpenAIEmbeddings() {
	_, err := Con.Exec(`
		CREATE TABLE IF NOT EXISTS database_settings (
			key TEXT PRIMARY KEY,
			value JSONB NOT NULL)`)
	if err != nil {
		panic(err)
	}
}
