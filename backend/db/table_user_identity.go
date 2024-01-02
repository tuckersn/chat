package db

import (
	"database/sql"
	"errors"
	"regexp"
	"time"
)

var UserNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_]+$`)
var DisplayNameRegex = regexp.MustCompile(`^[a-zA-Z0-9_ ]+$`)
var EmailRegex = regexp.MustCompile(`^([a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+|)$`)

const METADATA_KEY_TOTP_REQUEST = "totp_request"

type RecordUserTotpRequestMetadata struct {
	Secret      string `db:"secret"`
	ImageData   string `db:"image_data"`
	RequestTime int64  `db:"request_time"`
}

type RecordUser struct {
	Id           int32      `db:"id"`
	Username     string     `db:"username"`
	DisplayName  string     `db:"display_name"`
	Email        *string    `db:"email"`
	JoinedTime   time.Time  `db:"joined_time"`
	LastSeen     *time.Time `db:"last_seen"`
	Activated    bool       `db:"activated"`
	Admin        bool       `db:"admin"`
	TotpSecret   []byte     `db:"totp_secret"`
	TotpVerified bool       `db:"totp_verified"`
	Metadata     string     `db:"metadata"`
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
	Exec(`
		CREATE TABLE IF NOT EXISTS user_identity (
			id SERIAL PRIMARY KEY,
			username TEXT NOT NULL,
			display_name TEXT NOT NULL,
			email TEXT DEFAULT '',
			activated BOOLEAN DEFAULT FALSE,
			admin BOOLEAN DEFAULT FALSE,
			joined_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			last_seen TIMESTAMP WITH TIME ZONE,
			totp_secret TEXT DEFAULT NULL,
			totp_verified BOOLEAN DEFAULT FALSE,
			metadata JSONB DEFAULT '{}'::JSONB
		);
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_username ON user_identity (username);
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_display_name ON user_identity (display_name);
	`)

	/**
	 * mv_user_identity_admin
	 */
	Exec(`
		CREATE MATERIALIZED VIEW IF NOT EXISTS mv_user_identity_admin AS
		SELECT * FROM user_identity WHERE admin = TRUE;
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS mv_idx_user_identity_admin_username ON mv_user_identity_admin (username);
	`)

	/**
	 * post setup steps
	 */
	user, err := InsertUser("admin", "Administrator", "", nil, nil)
	if err != nil {
		panic(err)
	}
	user.SetActive(true).SetAdmin(true)
}

func InsertUser(username string, displayName string, email string, totpSecret *string, metadataJson []byte) (RecordUser, error) {
	if !UserNameRegex.MatchString(username) {
		return RecordUser{}, errors.New("Invalid username")
	}

	if !DisplayNameRegex.MatchString(displayName) {
		return RecordUser{}, errors.New("Invalid display name")
	}

	if len(metadataJson) < 3 {
		metadataJson = []byte("{}")
	}

	input := map[string]interface{}{
		"username":     username,
		"display_name": displayName,
		"email":        email,
		"totp_secret":  totpSecret,
		"metadata":     metadataJson,
	}

	if totpSecret == nil {
		input["totp_secret"] = sql.NullString{}
	}

	rows, err := Con.NamedQuery(`
		INSERT INTO user_identity (username, display_name, email, totp_secret, metadata)
		VALUES (:username, :display_name, :email, :totp_secret, :metadata)
		RETURNING *
	`, input)
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

func GetTotpRequestMetadata(username string) (RecordUserTotpRequestMetadata, error) {
	var totpReq RecordUserTotpRequestMetadata
	err := Con.Get(&totpReq, `
		SELECT metadata#>>'{totp_request,Secret}' as secret,
			metadata#>>'{totp_request,ImageData}' as image_data,
			metadata#>>'{totp_request,RequestTime}' as request_time
		FROM user_identity 
		WHERE username = $1
		`,
		username)
	if err != nil {
		return RecordUserTotpRequestMetadata{}, err
	}
	return totpReq, nil
}
