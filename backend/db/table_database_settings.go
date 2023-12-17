package db

import (
	"fmt"

	"github.com/tuckersn/chatbackend/openai"
	"github.com/tuckersn/chatbackend/util"
)

func SetDatabaseSetting(key string, value any) {

	valueStr := fmt.Sprint(value)

	Con.MustExec(`
		INSERT INTO database_settings (key, value)
		VALUES ($1, $2)
		ON CONFLICT(key) DO UPDATE
		SET value = $2
	`, key, valueStr)
}

/*
table name is plural since this is basically a string key to json map as a table
*/
func TableInitDatabaseSettings(context TableInitContext) {
	Exec(`
		CREATE TABLE IF NOT EXISTS database_settings (
			key TEXT PRIMARY KEY,
			value JSONB NOT NULL,
			last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS idx_database_settings_key ON database_settings (key);
	`)

	SetDatabaseSetting("version", 1)

	if util.Config.Production {
		SetDatabaseSetting("environment", "\"production\"")
	} else {
		SetDatabaseSetting("environment", "\"development\"")
	}

	if openai.APIKey() != "" {
		SetDatabaseSetting("openai_enabled", "true")
	} else {
		SetDatabaseSetting("openai_enabled", "false")
	}

}

func DBUpdateDatabaseSettings() {

}
