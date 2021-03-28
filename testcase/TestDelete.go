package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T, repo gorepo.Repository) {
	var err error
	ids := []string{"1", "2", "3", "4"}
	for _, id := range ids {
		err = repo.Delete("gorepo", id)
		assert.NoError(t, err)
	}
	data := []map[string]interface{}{}
	var filter gorepo.Filter
	err = repo.Fetch("gorepo", &filter, &data)
	assert.NoError(t, err)
	assert.Empty(t, data)
}
