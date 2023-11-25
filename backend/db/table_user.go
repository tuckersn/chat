package db

import "regexp"

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)

type User struct {
	Id          int32  `db:"id"`
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	JoinedTime  int64  `db:"joined_time"`
	LastSeen    *int64 `db:"last_seen"`
}

func TableInitUser() {
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS user (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		display_name TEXT NOT NULL,
		joined_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		last_seen TIMESTAMP WITH TIME ZONE,
		metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
	);
	`)
	if err != nil {
		panic(err)
	}

	_, err = Con.Exec(`
		INSERT INTO user (id, username, display_name)
		VALUES (0, 'admin', 'Administrator')
		ON CONFLICT DO NOTHING;
	`)
	if err != nil {
		panic(err)
	}

}

func CreateUser(username string, displayName string) {

	if !UserNameRegex.MatchString(username) {
		panic("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		panic("Invalid display name")
	}

	//TODO: add word filter
}

func GetUser(id int32) *User {
	var user User
	err := Con.Get(&user, "SELECT * FROM user WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return &user
}

func GetUserByUserName(username string) *User {
	var user User
	err := Con.Get(&user, "SELECT * FROM user WHERE username = $1", username)
	if err != nil {
		panic(err)
	}
	return &user
}
