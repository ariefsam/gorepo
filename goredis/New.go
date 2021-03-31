package goredis

func New(host, port, password string, database int) (redis Goredis) {
	redis.Host = host
	redis.Port = port
	if redis.Port == "" {
		redis.Port = "6379"
	}
	redis.Password = password
	redis.Database = database
	return
}
