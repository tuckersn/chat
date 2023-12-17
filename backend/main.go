package main

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/openai"
	"github.com/tuckersn/chatbackend/util"
)

func main() {

	logger := log.New(os.Stdout, "[START]", log.LstdFlags|log.Lshortfile)

	logger.Println("Loading env file")
	err := godotenv.Load()
	if err != nil {
		logger.Println("Error loading .env file")
		panic(err)
	}

	logger.Println("Loading config")
	util.LoadConfigOnStartup()

	logger.Println("Storage dir: " + util.GetStorageDir(""))
	util.CreateStorageDirectoryIfNotExists()

	logger.Println("Initializing cron scheduler")
	cron := func() gocron.Scheduler {
		timezone, err := time.LoadLocation(util.Config.Timezone)
		if err != nil {
			panic(err)
		}
		return *gocron.NewScheduler(timezone)
	}()

	logger.Println("Initializing database connection")
	db.InitializeDatabaseConnection(&cron)

	var models []openai.ModelResponse
	models, err = openai.GetModels()
	logger.Println("Models:", models)

	cron.StartAsync()
	httpServer()
}
