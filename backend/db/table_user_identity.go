package db

import (
	"errors"
	"regexp"
	"time"
)

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)

type RecordUser struct {
	Id          int32      `db:"id"`
	Username    string     `db:"username"`
	DisplayName string     `db:"display_name"`
	Email       *string    `db:"email"`
	JoinedTime  time.Time  `db:"joined_time"`
	LastSeen    *time.Time `db:"last_seen"`
	Activated   bool       `db:"activated"`
	Admin       bool       `db:"admin"`
	Metadata    string     `db:"metadata"`
}

func (user *RecordUser) SetActive(active bool) *RecordUser {
	_, err := Con.Exec("UPDATE user_identity SET activated = $1 WHERE id = $2", active, user.Id)
	if err != nil {
		panic(err)
	}
	return user
}

func (user *RecordUser) SetAdmin(admin bool) *RecordUser {
	_, err := Con.Exec("UPDATE user_identity SET admin = $1 WHERE id = $2", admin, user.Id)
	if err != nil {
		panic(err)
	}
	return user
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
			email TEXT,
			activated BOOLEAN DEFAULT FALSE,
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
	user, err := InsertUser("admin", "Administrator")
	if err != nil {
		panic(err)
	}
	user.SetActive(true).SetAdmin(true)
}

func InsertUser(username string, displayName string) (RecordUser, error) {
	if !UserNameRegex.MatchString(username) {
		return RecordUser{}, errors.New("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		return RecordUser{}, errors.New("Invalid display name")
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
		return RecordUser{}, err
	}

	var user RecordUser
	for rows.Next() {
		err = rows.StructScan(&user)
		return user, err
	}
	return RecordUser{}, errors.New("Failed to insert user")
}

func GetUserById(id int32) (RecordUser, error) {
	var user RecordUser
	err := Con.Get(&user, "SELECT * FROM user_identity WHERE id = $1", id)
	if err != nil {
		return RecordUser{}, err
	}
	return user, nil
}

func GetUser(username string) (RecordUser, error) {
	var user RecordUser
	err := Con.Get(&user, "SELECT * FROM user_identity WHERE username = $1", username)
	if err != nil {
		return RecordUser{}, err
	}
	return user, nil
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
