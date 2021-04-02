package gomysql

func New(connectionString string, database string) (gomysql Gomysql) {
	gomysql.Connection = connectionString
	return
}
