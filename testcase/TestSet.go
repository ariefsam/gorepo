package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestSet(t *testing.T, repo gorepo.Repository) {
	data := map[string]interface{}{
		"a": "b",
		"c": float64(10),
	}
	err := repo.Set("gorepo", "1", data)
	assert.NoError(t, err)

	getData := map[string]interface{}{}
	err = repo.Get("gorepo", "1", &getData)
	assert.NoError(t, err)

	for key, val := range data {
		assert.Equal(t, val, getData[key])
	}

	TestFetch(t, repo)
}
