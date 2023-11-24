package db

import "regexp"

var PagePathRegex = regexp.MustCompile(`^/([a-zA-Z0-9_\-]+)$`)

func DBInitializePageTable() {
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS page (
		id SERIAL PRIMARY KEY,
		owner_id INTEGER NOT NULL,
		path TEXT PRIMARY KEY,
		content TEXT NOT NULL,
		hash TEXT NOT NULL,
	);
	`)
	if err != nil {
		panic(err)
	}

	_, err = Con.Exec(`
	CREATE TABLE IF NOT EXISTS page_member (
		id SERIAL PRIMARY KEY,
		page_id INTEGER NOT NULL,
		user_id INTEGER NOT NULL,
		joined TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		FOREIGN KEY (page_id) REFERENCES page(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		panic(err)
	}
}
