package openai

import "os"

func APIKey() string {
	return os.Getenv("CR_PG_PGVECTOR_ENABLED")
}
