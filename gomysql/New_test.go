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
	ID string  `gorm:"primaryKey;column:id"`
	Ab string  `gorm:"column:a_B"`
	C  float64 `gorm:"column:c"`
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

	repo := gomysql.New(connectionString, "gorepo", "id")

	var abc Abc
	repo.Automigrate(&abc)
	repo.Model = abc

	testcase.TestSet(t, repo)
}
