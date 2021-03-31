package gorepo

type Redis interface {
	GetInt(key string) (data int, err error)
	GetString(key string) (data string, err error)
	Set(key string, data interface{}) (err error)
	SetEx(key string, data interface{}, expireSecond int) (err error)
	PFAdd(key string, data ...interface{}) (result int64, err error)
	PFCount(key ...string) (result int64, err error)
}
