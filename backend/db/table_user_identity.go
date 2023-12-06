package db

import (
	"errors"
	"regexp"
	"time"
)

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)

type User struct {
	Id          int32      `db:"id"`
	Key         string     `db:"key"`
	Username    string     `db:"username"`
	DisplayName string     `db:"display_name"`
	JoinedTime  time.Time  `db:"joined_time"`
	LastSeen    *time.Time `db:"last_seen"`
	Admin       bool       `db:"admin"`
	Metadata    string     `db:"metadata"`
}

func TableInitUserIdentity(context TableInitContext) {

	/**
	 * user_identity
	 */
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS user_identity (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			display_name TEXT NOT NULL,
			admin BOOLEAN DEFAULT FALSE,
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

	/**
	 * mv_user_identity_admin
	 */
	Con.MustExec(`
		CREATE MATERIALIZED VIEW IF NOT EXISTS mv_user_identity_admin AS
		SELECT * FROM user_identity WHERE admin = TRUE;
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS mv_idx_user_identity_admin_username ON mv_user_identity_admin (username);
	`)

	/**
	 * post setup steps
	 */
	InsertUser("admin", "Administrator")
	MakeUserAdmin("admin")
}

func InsertUser(username string, displayName string) (*User, error) {
	if !UserNameRegex.MatchString(username) {
		return nil, errors.New("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		return nil, errors.New("Invalid display name")
	}

	rows, err := Con.NamedQuery(`
		INSERT INTO user_identity (username, display_name)
		VALUES (:username, :display_name)
		RETURNING *
	`, map[string]interface{}{
		"username":     username,
		"display_name": displayName,
	})

	if err != nil {
		return nil, err
	}

	var user User
	for rows.Next() {
		err = rows.StructScan(&user)
		return &user, err
	}
	return nil, errors.New("Failed to insert user")
}

func InsertAdmin(username string, displayName string) {
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

func GetUserById(id int32) *User {
	var user User
	err := Con.Get(&user, "SELECT * FROM user_identity WHERE id = $1", id)
	if err != nil {
		panic(err)
	}
	return &user
}

func GetUser(username string) (*User, error) {
	var user User
	err := Con.Get(&user, "SELECT * FROM user_identity WHERE username = $1", username)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func MakeUserAdmin(username string) {
	_, err := Con.Exec("UPDATE user_identity SET admin = TRUE WHERE username = $1", username)
	if err != nil {
		panic(err)
	}
}

func IsAdmin(username string) bool {
	var count int32
	err := Con.Get(&count, "SELECT COUNT(*) FROM user_identity WHERE username = $1 AND admin = TRUE", username)
	if err != nil {
		panic(err)
	}
	return count > 0
}

func DeleteUser(username string) error {
	_, err := Con.Exec("DELETE FROM user_identity WHERE username = $1", username)
	return err
}

func DeleteUserById(id int32) error {
	_, err := Con.Exec("DELETE FROM user_identity WHERE id = $1", id)
	return err
}
