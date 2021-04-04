package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T, repo gorepo.Repository) {
	data := map[string]interface{}{
		"id": "1",
		"a":  "b",
		"c":  float64(10),
	}
	err := repo.Create(data)
	assert.NoError(t, err)

	getData := map[string]interface{}{}
	err = repo.Get("1", &getData)
	assert.NoError(t, err)

	for key, val := range data {
		assert.Equal(t, val, getData[key])
	}

	data["a"] = "b edited"

	err = repo.Update("1", data)
	assert.NoError(t, err)

	getData = map[string]interface{}{}
	err = repo.Get("1", &getData)
	assert.NoError(t, err)

	for key, val := range data {
		assert.Equal(t, val, getData[key])
	}

	TestFetch(t, repo)
}
