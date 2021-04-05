package gomysql

func New(connectionString string, tableName string, primaryKey string) (gomysql Gomysql) {
	gomysql.Connection = connectionString
	gomysql.TableName = tableName
	gomysql.PrimaryKey = primaryKey
	return
}
