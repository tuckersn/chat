package db

func TableInitMessageAttachment() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS message_attachment (
			id SERIAL PRIMARY KEY,
			message_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			room_id INTEGER NOT NULL,
			name TEXT NOT NULL,
			mime_type TEXT NOT NULL,
			size INTEGER NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			path TEXT NOT NULL,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (message_id) REFERENCES message(id),
			FOREIGN KEY (user_id) REFERENCES user_identity(id),
			FOREIGN KEY (room_id) REFERENCES room(id)
		);
	`)
}
