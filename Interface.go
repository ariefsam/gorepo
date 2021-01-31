package gorepo

type Interface interface {
	Set(tableName string, id string, data map[string]interface{}) (err error)
	Get(tableName string, id string) (data map[string]interface{}, err error)
}
