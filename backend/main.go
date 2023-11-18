package main

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func main() {

	db := func() *sqlx.DB {
		username := os.Getenv("CHATROOM_POSTGRES_USER")
		if username == "" {
			username = "postgres"
		}
		password := os.Getenv("CHATROOM_POSTGRES_PASSWORD")
		if password == "" {
			password = "postgres"
		}
		database := os.Getenv("CHATROOM_POSTGRES_DATABASE")
		if database == "" {
			database = "postgres"
		}
		host := os.Getenv("CHATROOM_POSTGRES_HOST")
		if host == "" {
			host = "localhost"
		}
		port := os.Getenv("CHATROOM_POSTGRES_PORT")
		if port == "" {
			port = "5432"
		}
		sslmode := os.Getenv("CHATROOM_POSTGRES_SSLMODE")
		if sslmode == "" {
			sslmode = "enable"
		}

		db, err := sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=%s", username, password, database, host, port, sslmode))
		if err != nil {
			panic(err)
		}
		return db
	}()

	httpServer(db)
}
