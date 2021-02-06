package gorepo

type Repository interface {
	Set(tableName string, id string, data interface{}) (err error)
	Get(tableName string, id string, result interface{}) (err error)
	Delete(tableName string, id string) (err error)
	Fetch(tableName string, filter map[string]interface{}, result interface{}) (err error)
	Sync(name string, to Repository)
	StopSync(name string)
	Export(name string, to Repository)
}
