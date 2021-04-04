package gomongo

import (
	"context"
	"log"
	"time"

	"github.com/ariefsam/gorepo"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2/bson"
)

type Gomongo struct {
	Connection     string
	Database       string
	PrimaryKey     string
	CollectionName string
}

const errorConnect = "Failed to connect mongodb"

func (gomongo Gomongo) primaryKey() (primaryKey string) {
	primaryKey = gomongo.PrimaryKey
	if primaryKey == "" {
		primaryKey = "_id"
	}
	return
}

func (gomongo Gomongo) Create(data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, errorConnect)
		return
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database(gomongo.Database).Collection(gomongo.CollectionName)
	primaryKey := gomongo.primaryKey()
	if primaryKey != "_id" {
		var opt options.IndexOptions
		t := true
		opt.Unique = &t
		mod := mongo.IndexModel{
			Keys: bson.M{
				primaryKey: 1, // index in ascending order
			}, Options: &opt,
		}
		s, err := coll.Indexes().CreateOne(context.TODO(), mod)
		if err != nil {
			log.Println(err)
		}
		log.Println(s)
	}

	_, err = coll.InsertOne(ctx, data)

	return
}

func (gomongo Gomongo) Update(id string, data interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, errorConnect)
		return
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database(gomongo.Database).Collection(gomongo.CollectionName)
	var option options.UpdateOptions
	t := true
	option.Upsert = &t
	primaryKey := gomongo.primaryKey()
	filter := bson.M{primaryKey: id}
	toUpdate := bson.M{"$set": data}

	_, err = coll.UpdateOne(ctx, filter, toUpdate, &option)

	return
}

func (gomongo Gomongo) Get(id string, result interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, errorConnect)
		return
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database(gomongo.Database).Collection(gomongo.CollectionName)
	primaryKey := gomongo.primaryKey()
	filter := bson.M{primaryKey: id}
	err = coll.FindOne(ctx, filter).Decode(result)
	if err != nil {
		return errors.New("Failed to find one in mongodb")
	}
	return
}
func (gomongo Gomongo) Fetch(filter *gorepo.Filter, result interface{}) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, errorConnect)
		return
	}
	defer client.Disconnect(context.TODO())
	coll := client.Database(gomongo.Database).Collection(gomongo.CollectionName)

	var filterMongo map[string]interface{}
	option := options.Find()
	if filter != nil {
		if filter.Where != nil {
			filterMongo = filter.Where
		}

		if filter.Sort != nil {
			option.SetSort(filter.Sort)
		}

		if filter.Limit != 0 {
			option.SetLimit(int64(filter.Limit))
		}
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
func (gomongo Gomongo) Delete(id string) (err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(gomongo.Connection))
	if err != nil {
		err = errors.Wrap(err, errorConnect)
		return
	}
	defer client.Disconnect(context.TODO())

	coll := client.Database(gomongo.Database).Collection(gomongo.CollectionName)
	primaryKey := gomongo.primaryKey()
	filter := bson.M{primaryKey: id}

	_, err = coll.DeleteOne(ctx, filter)

	return
}
