package db

type RecordUserIdentityGoogleRequest struct {
	Id        int    `db:"id"`
	MadeOn    string `db:"made_on"`
	CsrfToken string `db:"csrf_token"`
	OriginUrl string `db:"origin_url"`
}

func TableInitUserIdentityGoogleRequests(context TableInitContext) {
	Exec(`
		CREATE TABLE IF NOT EXISTS user_identity_google_requests (
			id SERIAL PRIMARY KEY,
			made_on TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			csrf_token TEXT NOT NULL,
			origin_url TEXT NOT NULL
		);
	`)

	Exec(`
		CREATE INDEX IF NOT EXISTS idx_user_identity_google_requests ON user_identity_google_requests (csrf_token);
	`)

	context.Cron.Every(10).Minutes().Do(func() {
		Exec(`
			DELETE FROM user_identity_google_requests
			WHERE made_on < NOW() - INTERVAL '10 minutes'
		`)
	})

}

func InsertUserIdentityGoogleRequest(csrf_token string, origin_url string) error {
	_, err := Con.NamedExec(`
		INSERT INTO user_identity_google_requests (csrf_token, origin_url)
		VALUES (:csrf_token, :origin_url)
	`, map[string]interface{}{
		"csrf_token": csrf_token,
		"origin_url": origin_url,
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
