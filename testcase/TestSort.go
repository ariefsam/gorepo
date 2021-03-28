package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T, repo gorepo.Repository) {
	sort := map[string]interface{}{
		"c": -1,
	}
	getData := []map[string]interface{}{}
	var filter gorepo.Filter
	filter.Sort = sort
	err := repo.Fetch("gorepo", filter, &getData)
	assert.NoError(t, err)
	assert.NotEmpty(t, getData)
	if len(getData) > 0 {
		assert.Equal(t, "4", getData[0]["id"])
	}

	sort = map[string]interface{}{
		"c": 1,
	}
	getData = []map[string]interface{}{}
	filter.Sort = sort
	err = repo.Fetch("gorepo", filter, &getData)
	assert.NoError(t, err)
	assert.NotEmpty(t, getData)
	if len(getData) > 0 {
		assert.Equal(t, "1", getData[0]["id"])
	}
	TestDelete(t, repo)
}
