package db

import "regexp"

var WebhookNameRegexp = regexp.MustCompile(`^[a-zA-Z0-9_\-]+$`)
var WebhookUrlRegexp = regexp.MustCompile(`^https?://[a-zA-Z0-9_\-\.]+(:[0-9]+|)?(/[a-zA-Z0-9_\-\.]+)*$`)

type RecordWebhook struct {
	Id       int32  `db:"id"`
	Name     string `db:"name"`
	Url      string `db:"url"`
	OwnerId  int32  `db:"owner_id"`
	Headers  string `db:"headers"`
	Subjects string `db:"subjects"`
	Metadata string `db:"metadata"`
	Created  int64  `db:"created"`
}

func TableInitWebhook(context TableInitContext) {
	Con.MustExec(`
		CREATE TABLE IF NOT EXISTS webhook (
			id SERIAL PRIMARY KEY,
			name TEXT NOT NULL,
			url TEXT NOT NULL,
			owner_id INTEGER NOT NULL,
			headers JSONB NOT NULL DEFAULT '{}'::JSONB,
			subjects JSONB NOT NULL DEFAULT '{}'::JSONB,
			metadata JSONB NOT NULL DEFAULT '{}'::JSONB,
			created TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (owner_id) REFERENCES user_identity(id)
		);
	`)

	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_webhook_owner_id ON webhook (owner_id);`)
	Con.MustExec(`CREATE INDEX IF NOT EXISTS idx_webhook_name ON webhook (name);`)
}

func InsertWebhook(name string, url string, owner_id int32) {
	if !WebhookNameRegexp.MatchString(name) {
		panic("Invalid webhook name")
	}

	if !WebhookUrlRegexp.MatchString(url) {
		panic("Invalid webhook url")
	}

	_, err := Con.NamedExec(`
		INSERT INTO webhook (name, url, owner_id)
		VALUES (:name, :url, :owner_id)
	`, map[string]interface{}{
		"name":     name,
		"url":      url,
		"owner_id": owner_id,
	})

	if err != nil {
		panic(err)
	}
}

func GetWebhook(id int32) *RecordWebhook {
	var webhook RecordWebhook
	err := Con.Get(&webhook, `
		SELECT * FROM webhook WHERE id = $1
	`, id)
	if err != nil {
		panic(err)
	}
	return &webhook
}
