package mongo

import "github.com/ariefsam/gorepo"

type Gomongo struct {
	Connection string
	Database   string
}

func (gomongo *Gomongo) Set(tableName string, id string, data interface{}) (err error) {
	return
}
func (gomongo *Gomongo) Get(tableName string, id string, result interface{}) (err error) {
	return
}
func (gomongo *Gomongo) Fetch(tableName string, filter map[string]interface{}, result interface{}) (err error) {
	return
}
func (gomongo *Gomongo) Sync(name string, to gorepo.Repository) {
	return
}
func (gomongo *Gomongo) StopSync(name string) {
	return
}
func (gomongo *Gomongo) Export(name string, filter map[string]interface{}, to gorepo.Repository) {
	return
}
