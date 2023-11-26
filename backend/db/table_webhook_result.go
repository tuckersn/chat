package db

func TableInitWebhookResult() {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS webhook_result (
			id SERIAL PRIMARY KEY,
			webhook_id INTEGER NOT NULL,
			trigerred_by INTEGER NOT NULL,
			subject TEXT NOT NULL,
			status_code INTEGER NOT NULL,
			status_text TEXT NOT NULL,
			request_time TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			FOREIGN KEY (webhook_id) REFERENCES webhook(id),
			FOREIGN KEY (trigerred_by) REFERENCES user_identity(id)
		);
	`)
}

func InsertWebhookResult(
	webhook_id int32,
	trigerred_by int32,
	subject string,
	status_code int32,
	status_text string,
) (any, error) {
	_, err := Con.NamedExec(`
		INSERT INTO webhook_result (webhook_id, trigerred_by, subject, status_code, status_text)
		VALUES (:webhook_id, :trigerred_by, :subject, :status_code, :status_text)
	`, map[string]interface{}{
		"webhook_id":   webhook_id,
		"trigerred_by": trigerred_by,
		"subject":      subject,
		"status_code":  status_code,
		"status_text":  status_text,
	})
	if err != nil {
		return nil, err
	}
	return nil, nil
}
