package util

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/tuckersn/chatbackend/db"
)

func TestSetup(relativeRootPath string) (*log.Logger, error) {
	var logger = log.New(os.Stdout, "[TEST]", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load(relativeRootPath)
	if err != nil {
		return nil, err
	}
	db.InitializeDatabaseConnection()
	return logger, nil
}
