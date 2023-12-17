package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-co-op/gocron"
	"github.com/jmoiron/sqlx"
	"github.com/tuckersn/chatbackend/openai"
	"github.com/tuckersn/chatbackend/util"
)

var Con *sqlx.DB = nil

type TableSharedContext struct {
	VectorsEnabled bool
	OpenAIEnabled  bool
}

type Table struct {
	Name string
	Init func(args TableInitContext)
}

type TableInitContext struct {
	Name          string
	Cron          *gocron.Scheduler
	SharedContext *TableSharedContext
}

// In-order of initialization
var Tables = []Table{
	{"databast_settings", TableInitDatabaseSettings},
	{"user_identity", TableInitUserIdentity},
	{"room", TableInitRoom},
	{"message", TableInitMessage},
	{"message_attachment", TableInitMessageAttachment},
	{"note", TableInitNote},
	{"user_identity_google", TableInitUserIdentityGoogle},
	{"user_identity_google_requests", TableInitUserIdentityGoogleRequests},
	{"login", TableInitLogin},
	{"webhook", TableInitWebhook},
	{"webhook_result", TableInitWebhookResult},
}

var databaseInitialized = false

func InitializeDatabaseConnection(cron *gocron.Scheduler) {

	if databaseInitialized {
		return
	}
	databaseInitialized = true

	log := log.New(os.Stdout, "[START][DB]", log.LstdFlags|log.Lshortfile)
	dbConf := util.Config.Database

	log.Printf("Connecting to database %s:%d", dbConf.Host, dbConf.Port)
	var err error
	Con, err = sqlx.Connect("postgres",
		fmt.Sprintf("user=%s password=%s database=%s host=%s port=%d sslmode=%s",
			dbConf.Username,
			dbConf.Password,
			dbConf.Database,
			dbConf.Host,
			dbConf.Port,
			dbConf.SSLMode))
	if err != nil {
		panic(err)
	}

	log.Printf("Checking for %s schema", dbConf.Schema)

	var schemaName string
	err = Con.Get(&schemaName, `
		SELECT schema_name
		FROM information_schema.schemata
		WHERE schema_name = $1;
		`, dbConf.Schema)

	if err != nil {
		if err == sql.ErrNoRows {
			log.Println(err)
		} else {
			log.Fatalln(err)
		}
	}

	if schemaName != dbConf.Schema {
		log.Printf("Creating schema %s", dbConf.Schema)
		_, err = Con.Exec(fmt.Sprintf("CREATE SCHEMA %s", dbConf.Schema))
		if err != nil {
			log.Fatalln(err)
		}
	}

	log.Printf("SET search_path TO %s", dbConf.Schema)
	_, err = Con.Exec(fmt.Sprintf("SET search_path TO %s", dbConf.Schema))
	if err != nil {
		log.Fatalln(err)
	}

	var shared TableSharedContext = TableSharedContext{
		VectorsEnabled: util.Config.Database.PGVector.Enabled,
		OpenAIEnabled:  openai.APIKey() != "",
	}
	if shared.VectorsEnabled {
		log.Println("pgvector is enabled, enabling vector similarity search and vector fields")
		_, err = Con.Exec(`
			CREATE EXTENSION IF NOT EXISTS vector CASCADE;
		`)
		if err != nil {
			log.Println("Error trying to enable pgvector extension")
			log.Println(err)
		}
		if shared.OpenAIEnabled {
			log.Println("OpenAI API key is set, enabling OpenAI embeddings (vector fields)")
			TableInitOpenAIEmbeddings(TableInitContext{
				Name:          "openai_embeddings",
				Cron:          cron,
				SharedContext: &shared,
			})
		}
	}

	for i, table := range Tables {
		log.Printf("[%d/%d] Initializing table %s", i+1, len(Tables), table.Name)
		table.Init(TableInitContext{
			Name:          table.Name,
			Cron:          cron,
			SharedContext: &shared,
		})
	}

}

func Exec(query string) {
	_, err := Con.Exec(query)
	if err != nil {
		panic(err)
	}
}

func TypeExists(typeName string) bool {
	var name string
	err := Con.Get(&name, `
		SELECT typname
		FROM pg_type
		WHERE typname = $1;
		`, typeName)
	if err != nil {
		if err == sql.ErrNoRows {
			return false
		}
		panic(err)
	}
	return name == typeName
}
