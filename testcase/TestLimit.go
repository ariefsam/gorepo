package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestLimit(t *testing.T, repo gorepo.Repository) {
	var err error

	data := []map[string]interface{}{}
	var filter gorepo.Filter
	filter.Limit = 2
	err = repo.Fetch("gorepo", &filter, &data)
	assert.NoError(t, err)
	assert.NotEmpty(t, data)
	assert.Len(t, data, 2)
	TestDelete(t, repo)
}
