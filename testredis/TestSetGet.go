package testredis

import (
	"testing"
	"time"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestSetGet(t *testing.T, redis gorepo.Redis) {
	var expected int
	expected = 5
	err := redis.Set("a", expected)
	assert.NoError(t, err)
	data, err := redis.GetInt("a")
	assert.NoError(t, err)
	assert.Equal(t, 5, data)

	expectedString := "abc"
	err = redis.SetEx("b", expectedString, 1)
	assert.NoError(t, err)
	dataString, err := redis.GetString("b")
	assert.NoError(t, err)
	assert.Equal(t, expectedString, dataString)
	time.Sleep(1 * time.Second)

	dataString, err = redis.GetString("b")
	assert.Error(t, err)
	assert.Equal(t, "", dataString)

}
