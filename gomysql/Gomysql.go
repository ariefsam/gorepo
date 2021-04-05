package gomysql

import (
	"github.com/ariefsam/gorepo"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Gomysql struct {
	Connection string
	TableName  string
	PrimaryKey string
	Model      interface{}
}

func (g Gomysql) primaryKey() (primaryKey string) {
	primaryKey = g.PrimaryKey
	if primaryKey == "" {
		primaryKey = "id"
	}
	return
}

const errorConnect = "Failed to connect mysql"

func (g Gomysql) Automigrate(data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}
	db.AutoMigrate(data)
	return
}

func (g Gomysql) Create(data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.Table(g.TableName).Create(data).Error

	return
}

func (g Gomysql) Update(id string, data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}
	where := map[string]interface{}{
		g.primaryKey(): id,
	}
	err = db.Table(g.TableName).Where(where).Updates(data).Error
	return
}

func (g Gomysql) Get(id string, data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}
	where := map[string]interface{}{
		g.primaryKey(): id,
	}

	err = db.Table(g.TableName).Where(where).Scan(data).Error
	return
}

func (g Gomysql) Fetch(filter *gorepo.Filter, data interface{}) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}
	dbq := db.Table(g.TableName)
	if filter != nil {

		if filter.Where != nil {
			dbq = dbq.Where(filter.Where)
		}

		if filter.Sort != nil {
			for key, val := range filter.Sort {
				order := "asc"
				if v, ok := val.(int); ok {
					if v == -1 {
						order = "desc"
					}
				}
				dbq = dbq.Order(key + " " + order)
			}
		}

		if filter.Limit != 0 {
			dbq = dbq.Limit(filter.Limit)
		}

	}

	err = dbq.Find(data).Error

	return
}

func (g Gomysql) Delete(id string) (err error) {
	db, err := gorm.Open(mysql.Open(g.Connection), &gorm.Config{})
	if err != nil {
		return
	}

	err = db.Delete(g.Model, id).Error
	return
}
