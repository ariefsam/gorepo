package gorepo

type Repository interface {
	Create(data interface{}) (err error)
	Update(id string, data interface{}) (err error)
	Get(id string, result interface{}) (err error)
	Delete(id string) (err error)
	Fetch(filter *Filter, result interface{}) (err error)
}
