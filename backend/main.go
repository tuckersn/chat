package main

import (
	"fmt"
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

	var logger = log.New(os.Stdout, "[START]", log.LstdFlags|log.Lshortfile)

	logger.Println("Loading env file")
	err := godotenv.Load()
	if err != nil {
		logger.Println("Error loading .env file")
		panic(err)
	}

	util.LoadConfigOnStartup()

	fmt.Println("Storage dir: " + util.GetStorageDir(""))
	util.CreateStorageDirectoryIfNotExists()

	s := gocron.NewScheduler(time.UTC)

	logger.Println("Initializing database connection")
	db.InitializeDatabaseConnection(s)

	var models []openai.ModelResponse
	models, err = openai.GetModels()
	fmt.Println(models)

	httpServer()
}
