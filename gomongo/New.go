package gomongo

func New(connectionString string, database string, collectionName string, primaryKey string) (gomongo Gomongo) {
	gomongo.Connection = connectionString
	gomongo.Database = database
	gomongo.CollectionName = collectionName
	gomongo.PrimaryKey = primaryKey
	return
}
