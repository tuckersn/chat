package db

import "github.com/tuckersn/chatbackend/util"

const MESSAGE_KEY_LENGTH = 10

type RecordMessage struct {
	Id       int32  `db:"id"`
	Key      string `db:"key"`
	RoomId   int32  `db:"room_id"`
	AuthorId int32  `db:"author_id"`
	Content  string `db:"content"`
	Created  int64  `db:"created"`
	Metadata string `db:"metadata"`
}

func (m *RecordMessage) Author() *User {
	return GetUserById(m.AuthorId)
}

func TableInitMessage(context TableInitContext) {
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

	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_message_room_id ON message (room_id);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_message_author_id ON message (author_id);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_message_key ON message (key);`)
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

func GetMessageById(id int32) *RecordMessage {
	var message RecordMessage
	err := Con.Get(&message, `
		SELECT * FROM message WHERE id = $1
	`, id)
	if err != nil {
		panic(err)
	}
	return &message

}

func GetMessage(key string) *RecordMessage {
	var message RecordMessage
	err := Con.Get(&message, `
		SELECT * FROM message WHERE key = $1
	`, key)
	if err != nil {
		panic(err)
	}
	return &message
}

func DeleteMessage(key string) {
	_, err := Con.Exec(`
		DELETE FROM message WHERE key = $1
	`, key)
	if err != nil {
		panic(err)
	}
}

func DeleteMessageById(id int32) {
	_, err := Con.Exec(`
		DELETE FROM message WHERE id = $1
	`, id)
	if err != nil {
		panic(err)
	}
}

func InsertNewMessage(room_key string, author_id int32, content string) *RecordMessage {
	var message RecordMessage
	key := util.RandomString(MESSAGE_KEY_LENGTH, []string{util.ALPHABET_ALL})
	err := Con.Get(&message, `
		INSERT INTO message (key, room_id, author_id, content)
		VALUES ($1, $2, $3, $4)
		RETURNING *
	`, key, room_key, author_id, content)
	if err != nil {
		panic(err)
	}
	return &message
}
