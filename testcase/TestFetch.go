package testcase

import (
	"testing"

	"github.com/ariefsam/gorepo"
	"github.com/stretchr/testify/assert"
)

func TestFetch(t *testing.T, repo gorepo.Repository) {
	data := map[string]interface{}{
		"id":  "2",
		"a_B": "c",
		"c":   float64(11),
	}
	err := repo.Create(data)
	assert.NoError(t, err)

	data2 := map[string]interface{}{
		"id":  "3",
		"a_B": "d",
		"c":   float64(12),
	}
	err = repo.Update("3", data2)
	assert.NoError(t, err)

	type Abc struct {
		ID string  `bson:"id"`
		Ab string  `bson:"a_B"`
		C  float64 `bson:"c"`
	}

	data3 := Abc{
		ID: "4",
		Ab: "d",
		C:  13,
	}
	err = repo.Create(data3)
	assert.NoError(t, err)

	var getAbc Abc
	err = repo.Get("4", &getAbc)
	assert.NoError(t, err)
	assert.Equal(t, data3, getAbc)

	getAbcMap := map[string]interface{}{}
	expectedAbcMap := map[string]interface{}{
		"id":  "4",
		"a_B": "d",
		"c":   float64(13),
	}
	err = repo.Get("4", &getAbcMap)
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
	err = repo.Fetch(&filter, &getData)
	assert.NoError(t, err)
	assert.NotEmpty(t, getData)
	if len(getData) > 0 {
		for key, val := range data {
			assert.Equal(t, val, getData[0][key])
		}
	}

	TestFetchAll(t, repo)
}
