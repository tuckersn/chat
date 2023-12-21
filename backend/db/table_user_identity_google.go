package db

import (
	"errors"
	"time"
)

type RecordUserIdentityGoogle struct {
	Id         int32     `db:"id"`
	UserId     int32     `db:"user_id"`
	GoogleId   string    `db:"google_id"`
	Active     bool      `db:"active"`
	LastUpdate time.Time `db:"last_updated"`
}

func TableInitUserIdentityGoogle(context TableInitContext) {
	Exec(`
		CREATE TABLE IF NOT EXISTS user_identity_google (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL REFERENCES "user_identity" (id) ON DELETE CASCADE,
			google_id TEXT NOT NULL,
			active BOOLEAN NOT NULL DEFAULT TRUE,
			last_updated TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP
		);
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_google_user_id ON user_identity_google (user_id);
	`)
}

func InsertUserIdentityGoogle(user_id int32, google_id string) (RecordUserIdentityGoogle, error) {
	rows, err := Con.NamedQuery(`
		INSERT INTO user_identity_google (user_id, google_id)
		VALUES (:user_id, :google_id)
		RETURNING *
	`, map[string]interface{}{
		"user_id":   user_id,
		"google_id": google_id,
	})

	_, err = Con.NamedExec(`
		UPDATE user_identity
		SET metadata = jsonb_set(metadata, '{google_auth}', 'true', TRUE)
		WHERE id = :user_id
	`, map[string]interface{}{
		"user_id": user_id,
	})

	if err != nil {
		return RecordUserIdentityGoogle{}, err
	}

	for rows.Next() {
		var record RecordUserIdentityGoogle
		err = rows.StructScan(&record)
		if err != nil {
			return RecordUserIdentityGoogle{}, err
		}
		return record, nil
	}

	return RecordUserIdentityGoogle{}, errors.New("No rows returned")
}

func UpdateUserIdentityGoogle(user_id int32, google_id string) {
	_, err := Con.NamedExec(`
		UPDATE user_identity_google
		SET google_id = :google_id
		WHERE user_id = :user_id
	`, map[string]interface{}{
		"user_id":   user_id,
		"google_id": google_id,
	})

	if err != nil {
		panic(err)
	}
}

func GetUserIdentityGoogleByGoogleId(google_id string) (RecordUserIdentityGoogle, error) {
	var record RecordUserIdentityGoogle
	err := Con.Get(&record, `
		SELECT * FROM user_identity_google WHERE google_id = $1
	`, google_id)
	if err != nil {
		return RecordUserIdentityGoogle{}, err
	}
	return record, nil
}

func DeleteIdentityGoogleByUserId(user_id int32) error {
	_, err := Con.Exec(`
		DELETE FROM user_identity_google WHERE user_id = $1
	`, user_id)
	if err != nil {
		return err
	}
	return nil
}
