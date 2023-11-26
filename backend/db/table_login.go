package db

type RecordLogin struct {
	Id           int32  `json:"id"`
	UserId       int32  `json:"user_id"`
	SessionStart int64  `json:"session_start"`
	Ip           string `json:"ip"`
	Metadata     string `json:"metadata"`
}

func TableInitLogin() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS login (
			id SERIAL PRIMARY KEY,
			user_id INTEGER NOT NULL,
			session_start TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			ip TEXT NOT NULL,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (user_id) REFERENCES user_identity(id)
		);
	`)
}

func InsertLogin(user_id int32, ip string) (any, error) {
	_, err := Con.Exec(`
		INSERT INTO login (user_id, ip) VALUES ($1, $2);
	`, user_id, ip)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func SelectLastLoginByUserId(user_id int32) (any, error) {
	var records []RecordLogin
	err := Con.Select(&records, `
		SELECT * FROM login WHERE user_id = $1 ORDER BY session_start DESC LIMIT 1;
	`, user_id)
	if err != nil {
		return nil, err
	}
	return records, nil
}
