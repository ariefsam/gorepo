package gomongo_test

import (
	"log"
	"os"
	"testing"

	"github.com/ariefsam/gorepo/gomongo"
	"github.com/ariefsam/gorepo/testcase"
	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	databaseName := os.Getenv("MONGODB_DB_NAME")

	repo := gomongo.New(connectionString, databaseName, "gorepo", "id")
	repo.PrimaryKey = "id"

	testcase.TestSet(t, repo)
}
