package db

import "regexp"

var PagePathRegex = regexp.MustCompile(`^/([a-zA-Z0-9_\-]+)$`)

func TableInitPage() {
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS page (
		id SERIAL PRIMARY KEY,
		owner_id INTEGER NOT NULL,
		path TEXT PRIMARY KEY,
		content TEXT NOT NULL,
		created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
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
		metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
		FOREIGN KEY (page_id) REFERENCES page(id),
		FOREIGN KEY (user_id) REFERENCES users(id)
	);
	`)
	if err != nil {
		panic(err)
	}
}

func CreatePage(path string, content string, owner_id int32) {

	if !PagePathRegex.MatchString(path) {
		panic("Invalid page path")
	}

}
