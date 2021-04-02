package gomysql_test

import (
	"log"
	"os"
	"testing"

	"github.com/ariefsam/gorepo/gomysql"
	"github.com/ariefsam/gorepo/testcase"
	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connectionString := os.Getenv("MYSQL_CONNECTION_STRING")
	databaseName := os.Getenv("MYSQL_DB_NAME")

	repo := gomysql.New(connectionString, databaseName)

	testcase.TestSet(t, repo)
}
