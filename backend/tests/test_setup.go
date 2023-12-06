package tests

import (
	"log"
	"os"
	"time"

	"github.com/go-co-op/gocron"
	"github.com/joho/godotenv"
	"github.com/tuckersn/chatbackend/db"
)

func TestSetup(relativeRootPath string) (*log.Logger, error) {
	var logger = log.New(os.Stdout, "[TEST]", log.LstdFlags|log.Lshortfile)
	err := godotenv.Load(relativeRootPath)
	if err != nil {
		return nil, err
	}
	s := gocron.NewScheduler(time.UTC)
	db.InitializeDatabaseConnection(s)
	return logger, nil
}
