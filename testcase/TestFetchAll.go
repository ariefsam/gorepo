package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestFetchAll(t *testing.T, repo gorepo.Repository) {
	getData := []map[string]interface{}{}

	err := repo.Fetch(nil, &getData)
	assert.NoError(t, err)
	assert.Equal(t, 4, len(getData))
	TestSort(t, repo)
}
