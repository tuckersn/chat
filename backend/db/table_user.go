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

func DBInitializeUserTable() {
	_, err := Con.Exec(`
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username TEXT NOT NULL,
		display_name TEXT NOT NULL,
		joined_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
		last_seen TIMESTAMP WITH TIME ZONE
	);
	`)
	if err != nil {
		panic(err)
	}

	_, err = Con.Exec(`
		INSERT INTO users (id, username, display_name)
		VALUES (0, 'admin', 'Administrator')
		ON CONFLICT DO NOTHING;
	`)
	if err != nil {
		panic(err)
	}

}

func DBCreateUser(username string, displayName string) {

	if !UserNameRegex.MatchString(username) {
		panic("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		panic("Invalid display name")
	}

	//TODO: add word filter
}

func DBGetUser(id int32) *User {
	var user User
	err := Con.Get(&user, "SELECT * FROM users WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return &user
}

func DBGetUserByUsername(username string) *User {
	var user User
	err := Con.Get(&user, "SELECT * FROM users WHERE username = $1", username)
	if err != nil {
		panic(err)
	}
	return &user
}
