package db

type Message struct {
	Id       int32  `db:"id"`
	Key      string `db:"key"`
	RoomId   int32  `db:"room_id"`
	AuthorId int32  `db:"author_id"`
	Content  string `db:"content"`
	Created  int64  `db:"created"`
	Metadata string `db:"metadata"`
}

func (m *Message) Author() *User {
	return GetUser(m.AuthorId)
}

func TableInitMessage() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS message (
			id SERIAL PRIMARY KEY,
			key TEXT NOT NULL,
			room_id INTEGER NOT NULL,
			author_id INTEGER NOT NULL,
			content TEXT NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (author_id) REFERENCES user_identity(id),
			FOREIGN KEY (room_id) REFERENCES room(id)
		);
	`)
}

func InsertMessage(room_id int32, author_id int32, content string) {
	_, err := Con.NamedExec(`
		INSERT INTO message (room_id, author_id, content)
		VALUES (:room_id, :author_id, :content)
	`, map[string]interface{}{
		"room_id":   room_id,
		"author_id": author_id,
		"content":   content,
	})
	if err != nil {
		panic(err)
	}
}
