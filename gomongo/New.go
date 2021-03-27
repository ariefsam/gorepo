package gomongo

func New(connectionString string, database string) (gomongo Gomongo) {
	gomongo.Connection = connectionString
	gomongo.Database = database
	return
}
