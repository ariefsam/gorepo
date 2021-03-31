package goredis_test

import (
	"log"
	"os"
	"strconv"
	"testing"

	"github.com/ariefsam/gorepo/goredis"
	"github.com/ariefsam/gorepo/testredis"
	"github.com/joho/godotenv"
)

func TestNew(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file")
	}

	host := os.Getenv("REDIS_HOST")
	port := os.Getenv("REDIS_PORT")
	password := os.Getenv("REDIS_PASSWORD")
	db, _ := strconv.Atoi(os.Getenv("REDIS_DB"))

	redis := goredis.New(host, port, password, db)

	testredis.TestSetGet(t, redis)

}
