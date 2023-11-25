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
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS messages (
		id SERIAL PRIMARY KEY,
		key TEXT NOT NULL,
		room_id INTEGER NOT NULL,
		author_id INTEGER NOT NULL,
		content TEXT NOT NULL,
		created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
		FOREIGN KEY (author_id) REFERENCES users(id),
		FOREIGN KEY (room_id) REFERENCES rooms(id)
	);`)
	if err != nil {
		panic(err)
	}
}

func CreateMessage(username string, displayName string) {

	if !UserNameRegex.MatchString(username) {
		panic("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		panic("Invalid display name")
	}

	//TODO: add word filter
}
