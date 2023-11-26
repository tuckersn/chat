package db

import "regexp"

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)

type User struct {
	Id          int32  `db:"id"`
	Key         string `db:"key"`
	Username    string `db:"username"`
	DisplayName string `db:"display_name"`
	JoinedTime  int64  `db:"joined_time"`
	LastSeen    *int64 `db:"last_seen"`
}

func TableInitUser() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS user_identity (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			display_name TEXT NOT NULL,
			joined_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			last_seen TIMESTAMP WITH TIME ZONE,
			metadata JSONB DEFAULT '{}'::JSONB
		);
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_username ON user_identity (username);
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_display_name ON user_identity (display_name);
	`)

	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS user_identity_password (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			password TEXT NOT NULL,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB DEFAULT '{}'::JSONB,
			FOREIGN KEY (user_id) REFERENCES user_identity(id)
		);
	`)
}

func InsertUser(username string, displayName string) {
	if !UserNameRegex.MatchString(username) {
		panic("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		panic("Invalid display name")
	}

	_, err := Con.NamedExec(`
		INSERT INTO user_identity (username, display_name)
		VALUES (:username, :display_name)
	`, map[string]interface{}{
		"username":     username,
		"display_name": displayName,
	})

	if err != nil {
		panic(err)
	}
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
