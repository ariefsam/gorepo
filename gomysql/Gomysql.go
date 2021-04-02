package gomysql

import "github.com/ariefsam/gorepo"

type Gomysql struct {
	Connection string
	PrimaryKey string
}

func (g Gomysql) Set(tablename string, id string, data interface{}) (err error) {
	return
}

func (g Gomysql) Fetch(tablename string, filter *gorepo.Filter, data interface{}) (err error) {
	return
}

func (g Gomysql) Delete(tablename string, id string) (err error) {
	return
}
