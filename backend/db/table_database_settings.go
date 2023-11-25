package db

import "os"

func SetDatabaseSetting(key string, value any) {
	_, err := Con.Exec(`
		INSERT INTO database_settings (key, value)
		VALUES ($1, $2)
		ON CONFLICT DO UPDATE
		SET value = $2
	`, key, value)
	if err != nil {
		panic(err)
	}
}

func TableInitDatabaseSettings() {
	_, err := Con.Exec(`
		CREATE TABLE IF NOT EXISTS database_settings (
			key TEXT PRIMARY KEY,
			value JSONB NOT NULL),
			last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP`)
	if err != nil {
		panic(err)
	}

	SetDatabaseSetting("version", 1)
	SetDatabaseSetting("production", os.Getenv("PRODUCTION") != "")
	SetDatabaseSetting("openai_enabled", os.Getenv("OPENAI_API_KEY") != "")

}

func DBUpdateDatabaseSettings() {

}
