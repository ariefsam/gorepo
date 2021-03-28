package gomongo

import (
	"context"
	"time"

	"github.com/ariefsam/gorepo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Gomongo struct {
	Connection string
	Database   string
}

func (gomongo Gomongo) Set(tableName string, id string, data interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, "Failed to connect mongodb")
		return
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database(gomongo.Database).Collection(tableName)
	var option options.UpdateOptions
	t := true
	option.Upsert = &t
	filter := bson.M{"id": id}
	toUpdate := bson.M{"$set": data}

	_, err = coll.UpdateOne(ctx, filter, toUpdate, &option)

	return
}
func (gomongo Gomongo) Get(tableName string, id string, result interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, "Failed to connect mongodb")
		return
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database(gomongo.Database).Collection(tableName)
	filter := bson.M{"id": id}
	err = coll.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return errors.New("Failed to find one in mongodb")
	}
	return
}
func (gomongo Gomongo) Fetch(tableName string, filter gorepo.Filter, result interface{}) (err error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, "Failed to connect mongodb")
		return
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database(gomongo.Database).Collection(tableName)

	var filterMongo map[string]interface{}
	option := options.Find()

	if filter.Where != nil {
		filterMongo = filter.Where
	}

	if filter.Sort != nil {
		option.SetSort(filter.Sort)
	}

	cur, err := coll.Find(ctx, filterMongo, option)
	if err != nil {
		return errors.Wrap(err, "Failed to find operation")
	}
	defer cur.Close(context.TODO())

	err = cur.All(ctx, result)
	if err != nil {
		return errors.New("Failed to decode all")
	}
	return
}
func (gomongo Gomongo) Delete(tableName string, id string) (err error) {
	return
}
func (gomongo Gomongo) Sync(name string, to gorepo.Repository) {
	return
}
func (gomongo Gomongo) StopSync(name string) {
	return
}
