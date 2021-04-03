package gomysql

import (
	"github.com/ariefsam/gorepo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gomysql struct {
	Connection string
	PrimaryKey string
}

func (g Gomysql) Automigrate(tablename string, data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(data)
	return
}

func (g Gomysql) Set(tablename string, id string, data interface{}) (err error) {
	// db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	return
}

func (g Gomysql) Get(tablename string, id string, data interface{}) (err error) {
	return
}

func (g Gomysql) Fetch(tablename string, filter *gorepo.Filter, data interface{}) (err error) {
	return
}

func (g Gomysql) Delete(tablename string, id string) (err error) {
	return
}
