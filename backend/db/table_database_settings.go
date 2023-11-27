package db

import (
	"os"

	"github.com/tuckersn/chatbackend/openai"
)

func SetDatabaseSetting(key string, value any) {
	Con.MustExec(`
		INSERT INTO database_settings (key, value)
		VALUES ($1, $2)
		ON CONFLICT(key) DO UPDATE
		SET value = $2
	`, key, value)
}

/*
table name is plural since this is basically a string key to json map as a table
*/
func TableInitDatabaseSettings() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS database_settings (
			key TEXT PRIMARY KEY,
			value JSONB NOT NULL,
			last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_database_settings_key ON database_settings (key);
	`)

	SetDatabaseSetting("version", 1)

	if os.Getenv("PRODUCTION") != "" {
		SetDatabaseSetting("environment", os.Getenv("PRODUCTION"))
	} else {
		SetDatabaseSetting("environment", "development")
	}

	if openai.APIKey() != "" {
		SetDatabaseSetting("openai_enabled", "true")
	} else {
		SetDatabaseSetting("openai_enabled", "false")
	}

}

func DBUpdateDatabaseSettings() {

}
