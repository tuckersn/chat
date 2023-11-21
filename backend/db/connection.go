package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

var Con *sqlx.DB = nil

type Table struct {
	Name string
	Init func()
}

// In-order of initialization
var Tables = []Table{
	{"databast_settings", DBInitializeDatabaseSettings},
	{"users", DBInitializeUserTable},
	{"rooms", DBInitializeRoomTable},
	{"messages", DBInitializeMessageTable},
}

func InitializeDatabaseConnection() {

	var log = log.New(os.Stdout, "[START][DB]", log.LstdFlags|log.Lshortfile)

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
		database = "chatroom"
	}
	schema := os.Getenv("CHATROOM_POSTGRES_SCHEMA")
	if schema == "" {
		schema = "public"
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

	log.Printf("Connecting to database %s:%s", host, port)
	var err error
	Con, err = sqlx.Connect("postgres", fmt.Sprintf("user=%s password=%s database=%s host=%s port=%s sslmode=%s", username, password, database, host, port, sslmode))
	if err != nil {
		panic(err)
	}

	log.Printf("Checking for %s schema", schema)

	var schemaName string
	err = Con.Get(&schemaName, `
		SELECT schema_name
		FROM information_schema.schemata
		WHERE schema_name = $1;
		`, schema)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		} else {
			log.Fatalln(err)
		}
	}

	if schemaName != schema {
		log.Printf("Creating schema %s", schema)
		_, err = Con.Exec(fmt.Sprintf("CREATE SCHEMA %s", schema))
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("SET search_path TO %s", schema)
	_, err = Con.Exec(fmt.Sprintf("SET search_path TO %s", schema))
	if err != nil {
		log.Fatalln(err)
	}

	for i, table := range Tables {
		log.Printf("[%d/%d] Initializing table %s", i+1, len(Tables), table.Name)
		table.Init()
	}

}
