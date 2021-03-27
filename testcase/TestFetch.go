package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T, repo gorepo.Repository) {
	data := map[string]interface{}{
		"a": "c",
		"c": float64(11),
	}
	err := repo.Set("gorepo", "2", data)
	assert.NoError(t, err)

	data = map[string]interface{}{
		"a": "d",
		"c": float64(12),
	}
	err = repo.Set("gorepo", "3", data)
	assert.NoError(t, err)

	getData := []map[string]interface{}{}
	var filter gorepo.Filter
	filter.Where = map[string]interface{}{
		"a": "c",
	}
	err = repo.Fetch("gorepo", filter, &getData)
	assert.NoError(t, err)
	assert.NotEmpty(t, getData)

}