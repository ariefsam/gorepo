package gorepo

type Filter struct {
	Where map[string]interface{}
	Sort  map[string]interface{}
	Limit int
}
