package db

func TableInitUserIdentityGoogle() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS user_identity_google (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES "user_identity" (id) ON DELETE CASCADE,
			google_id TEXT NOT NULL,
			active BOOLEAN NOT NULL DEFAULT TRUE,
			last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_google_user_id ON user_identity_google (user_id);
	`)
}
