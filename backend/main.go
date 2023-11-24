package main

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/tuckersn/chatbackend/db"
	"github.com/tuckersn/chatbackend/openai"
)

func main() {

	var logger = log.New(os.Stdout, "[START]", log.LstdFlags|log.Lshortfile)

	logger.Println("Loading env file")
	err := godotenv.Load()
	if err != nil {
		logger.Fatal("Error loading .env file")
	}

	logger.Println("Initializing database connection")
	db.InitializeDatabaseConnection()

	var models []openai.ModelResponse
	models, err = openai.GetModels()
	fmt.Println(models)

	httpServer()
}
