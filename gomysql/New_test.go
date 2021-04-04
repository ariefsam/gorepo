package gomysql_test

import (
	"log"
	"os"
	"testing"

	"github.com/ariefsam/gorepo/gomysql"
	"github.com/ariefsam/gorepo/testcase"
	"github.com/joho/godotenv"
)

type Abc struct {
	Ab string  `bson:"a_B"`
	C  float64 `bson:"c"`
}

func (a Abc) TableName() string {
	return "gorepo"
}
func TestNew(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connectionString := os.Getenv("MYSQL_CONNECTION_STRING")
	databaseName := os.Getenv("MYSQL_DB_NAME")

	repo := gomysql.New(connectionString, databaseName)

	var abc Abc
	repo.Automigrate("gorepo", &abc)

	testcase.TestSet(t, repo)
}
