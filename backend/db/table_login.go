package db

import (
	"encoding/json"
	"time"
)

const (
	RECORD_LOGIN_ORIGIN_WEB     = 0
	RECORD_LOGIN_ORIGIN_MOBILE  = 1
	RECORD_LOGIN_ORIGIN_DESKTOP = 2
	RECORD_LOGIN_ORIGIN_OTHER   = 3
	RECORD_LOGIN_ORIGIN_API     = 4
)

type RecordLogin struct {
	Id           int32     `db:"id"`
	UserId       int32     `db:"user_id"`
	Origin       string    `db:"origin"`
	SessionStart time.Time `db:"session_start"`
	Ip           string    `db:"ip"`
	Metadata     []byte    `db:"metadata"`
}

func loginTypeFromInt(origin int32) string {
	switch origin {
	case RECORD_LOGIN_ORIGIN_WEB:
		return "web"
	case RECORD_LOGIN_ORIGIN_MOBILE:
		return "mobile"
	case RECORD_LOGIN_ORIGIN_DESKTOP:
		return "desktop"
	case RECORD_LOGIN_ORIGIN_OTHER:
		return "other"
	case RECORD_LOGIN_ORIGIN_API:
		return "api"
	}
	return "unknown"
}

func TableInitLogin(context TableInitContext) {

	if !TypeExists("enum_login_origin") {
		Exec(`CREATE TYPE enum_login_origin AS ENUM (
			'web',
			'mobile',
			'desktop',
			'other',
			'api'
		);`)
	}

	Exec(`
		CREATE TABLE IF NOT EXISTS login (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			origin enum_login_origin NOT NULL DEFAULT 'web',
			session_start TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			ip TEXT NOT NULL,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (user_id) REFERENCES user_identity(id)
		);
	`)

	Exec(`CREATE INDEX IF NOT EXISTS idx_login_user_id ON login (user_id);`)
	Exec(`CREATE INDEX IF NOT EXISTS idx_login_ip ON login (ip);`)
}

func InsertLogin(user_id int32, origin int32, ip string, metadata interface{}) (any, error) {
	metaStr := "{}"
	if metadata != nil {
		metaBytes, err := json.Marshal(metadata)
		if err != nil {
			return nil, err
		}
		metaStr = string(metaBytes)
	}
	_, err := Con.Exec(`
		INSERT INTO login (user_id, ip, origin, metadata) VALUES ($1, $2, $3, $4);
	`, user_id, ip, loginTypeFromInt(origin), metaStr)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func GetLoginsByUserId(user_id int32) (any, error) {
	var records []RecordLogin
	err := Con.Select(&records, `
		SELECT * FROM login WHERE user_id = $1 ORDER BY session_start DESC LIMIT 1;
	`, user_id)
	if err != nil {
		return nil, err
	}
	return records, nil
}

func GetLoginsRecent(user_id int32, amount int32) (any, error) {
	var records []RecordLogin
	err := Con.Select(&records, `
		SELECT * FROM login WHERE user_id = $1 ORDER BY session_start DESC LIMIT $2;
	`, user_id, amount)
	if err != nil {
		return nil, err
	}
	return records, nil
}
