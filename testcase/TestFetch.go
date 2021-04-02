package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T, repo gorepo.Repository) {
	data := map[string]interface{}{
		"a_B": "c",
		"c":   float64(11),
	}
	err := repo.Set("gorepo", "2", data)
	assert.NoError(t, err)

	data2 := map[string]interface{}{
		"a_B": "d",
		"c":   float64(12),
	}
	err = repo.Set("gorepo", "3", data2)
	assert.NoError(t, err)

	type Abc struct {
		Ab string  `bson:"a_B"`
		C  float64 `bson:"c"`
	}

	data3 := Abc{
		Ab: "d",
		C:  13,
	}
	err = repo.Set("gorepo", "4", data3)
	assert.NoError(t, err)

	var getAbc Abc
	err = repo.Get("gorepo", "4", &getAbc)
	assert.NoError(t, err)
	assert.Equal(t, data3, getAbc)

	getAbcMap := map[string]interface{}{}
	expectedAbcMap := map[string]interface{}{
		"a_B": "d",
		"c":   float64(13),
	}
	err = repo.Get("gorepo", "4", &getAbcMap)
	assert.NoError(t, err)
	if len(getAbcMap) > 0 {
		for key, val := range expectedAbcMap {
			assert.Equal(t, val, getAbcMap[key])
		}
	}

	getData := []map[string]interface{}{}
	var filter gorepo.Filter
	filter.Where = map[string]interface{}{
		"a_B": "c",
	}
	err = repo.Fetch("gorepo", &filter, &getData)
	assert.NoError(t, err)
	assert.NotEmpty(t, getData)
	if len(getData) > 0 {
		for key, val := range data {
			assert.Equal(t, val, getData[0][key])
		}
	}

	TestFetchAll(t, repo)
}
