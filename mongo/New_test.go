package mongo_test

import (
	"log"
	"os"
	"testing"

	"github.com/ariefsam/gorepo/mongo"

	"github.com/joho/godotenv"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")
	database := os.Getenv("MONGODB_DB_NAME")

	repo := mongo.New(connectionString, database)

	data := map[string]interface{}{
		"a": "b",
		"c": 1,
	}
	err = repo.Set("gorepo", "1", data)
	assert.NoError(t, err)
}
