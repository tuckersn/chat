package database

import "regexp"

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)

var schema = `
CREATE TABLE IF NOT EXISTS users (
	username TEXT PRIMARY KEY,
	display_name TEXT NOT NULL,
	joined_time INTEGER NOT NULL,
	last_seen INTEGER
);
`

type User struct {
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	JoinedTime  int64  `db:"joined_time"`
	LastSeen    *int64 `db:"last_seen"`
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

func DBInitializeUserTable() {

}
