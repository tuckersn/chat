package db

type RecordUserIdentityGoogleRequest struct {
	Id        int    `db:"id"`
	MadeOn    string `db:"made_on"`
	CsrfToken string `db:"csrf_token"`
}

func TableInitUserIdentityGoogleRequests(context TableInitContext) {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS user_identity_google_requests (
			id SERIAL PRIMARY KEY,
			made_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			csrf_token TEXT NOT NULL
		);
	`)

	Con.MustExec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_google_requests ON user_identity_google_requests (csrf_token);
	`)

	context.Cron.Every(10).Minutes().Do(func() {
		Con.MustExec(`
			DELETE FROM user_identity_google_requests
			WHERE made_on < NOW() - INTERVAL '10 minutes'
		`)
	})
}

func InsertUserIdentityGoogleRequest(csrf_token string) error {
	_, err := Con.NamedExec(`
		INSERT INTO user_identity_google_requests (csrf_token)
		VALUES (:csrf_token)
	`, map[string]interface{}{
		"csrf_token": csrf_token,
	})
	if err != nil {
		return err
	}
	return nil
}

func GetUserIdentityGoogleRequestDeleteAfter(csrf_token string) (RecordUserIdentityGoogleRequest, error) {
	var request RecordUserIdentityGoogleRequest
	err := Con.Get(&request, `
		SELECT * FROM user_identity_google_requests WHERE csrf_token = $1
	`, csrf_token)
	if err != nil {
		return RecordUserIdentityGoogleRequest{}, err
	}
	_, err = Con.Exec(`
		DELETE FROM user_identity_google_requests WHERE csrf_token = $1
	`, csrf_token)
	if err != nil {
		return RecordUserIdentityGoogleRequest{}, err
	}
	return request, nil
}
