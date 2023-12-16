package db

import "regexp"

var NotePathRegex = regexp.MustCompile(`^/([a-zA-Z0-9_\-]+)$`)

type RecordNote struct {
	Id      int32  `db:"id"`
	Path    string `db:"path"`
	OwnerId int32  `db:"owner_id"`
	Content string `db:"content"`
	Created int64  `db:"created"`
}

func TableInitNote(context TableInitContext) {
	Exec(`
		CREATE TABLE IF NOT EXISTS note (
			id SERIAL PRIMARY KEY,
			path TEXT NOT NULL,
			owner_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (owner_id) REFERENCES user_identity(id)
		);
	`)

	Exec(`CREATE INDEX IF NOT EXISTS idx_note_path ON note (path);`)
	Exec(`CREATE INDEX IF NOT EXISTS idx_note_owner_id ON note (owner_id);`)

	Exec(`
		CREATE TABLE IF NOT EXISTS note_member (
			id SERIAL PRIMARY KEY,
			note_id INTEGER NOT NULL,
			user_id INTEGER NOT NULL,
			joined TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (note_id) REFERENCES note(id),
			FOREIGN KEY (user_id) REFERENCES user_identity(id)
		);
	`)

	Exec(`CREATE INDEX IF NOT EXISTS idx_note_member_note_id ON note_member (note_id);`)
	Exec(`CREATE INDEX IF NOT EXISTS idx_note_member_user_id ON note_member (user_id);`)
}

func InsertNote(path string, content string, owner_id int32) {

	if !NotePathRegex.MatchString(path) {
		panic("Invalid page path")
	}

	_, err := Con.NamedExec(`
		INSERT INTO note (path, content, owner_id)
		VALUES (:path, :content, :owner_id)
	`, map[string]interface{}{
		"path":     path,
		"content":  content,
		"owner_id": owner_id,
	})

	if err != nil {
		panic(err)
	}

}

func GetNote(path string) *RecordNote {
	var note RecordNote
	err := Con.Get(&note, "SELECT * FROM note WHERE path = $1", path)
	if err != nil {
		panic(err)
	}
	return &note
}

func GetNoteById(id int32) *RecordNote {
	var note RecordNote
	err := Con.Get(&note, "SELECT * FROM note WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return &note
}
