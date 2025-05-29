package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"order-package/internal/infra/config"
)

type handleError func(ctx context.Context) error

type MongoCollection struct {
	opts     *options.ClientOptions
	mongoCfg config.Mongo
}

type Collection interface {
	Find(ctx context.Context) []interface{}
	Create(ctx context.Context, document interface{}) error
	Delete(ctx context.Context, filter interface{}) error
}
type CollectionConfig string
const (
	CollectionName CollectionConfig = "collection"
)

func NewMongoCollection(mongoConfig config.Database) *MongoCollection {
	mongoURI := fmt.Sprintf("mongodb://%s:%s@%s:%s/?authSource=admin",
		mongoConfig.Mongo.Username,
		mongoConfig.Mongo.Password,
		mongoConfig.Mongo.Host,
		mongoConfig.Mongo.Port)
	return &MongoCollection{
		opts:     options.Client().ApplyURI(mongoURI),
		mongoCfg: mongoConfig.Mongo,
	}
}

func getMongoCollectionVar(ctx context.Context) string {
	return ctx.Value(CollectionName).(string)
}

func isError(ctx context.Context, handler handleError) {
	func(ctx context.Context) {
		err := handler(ctx)
		if err != nil {
			panic(err)
		}
	}(ctx)
}

func (mc MongoCollection) getCollection(ctx context.Context) (*mongo.Client, *mongo.Collection) {
	client, err := mongo.Connect(ctx, mc.opts)
	if err != nil {
		fmt.Printf("Error %s connecting to database ", err.Error())
		panic(err)
	}
	collectionName := getMongoCollectionVar(ctx)

	return client, client.Database(mc.mongoCfg.Database).Collection(collectionName)
}
func (c MongoCollection) Delete(ctx context.Context, filter interface{}) error {
	client, cll := c.getCollection(ctx)
	defer isError(ctx, client.Disconnect)
	if _, err := cll.DeleteOne(ctx, filter); err != nil {
		return err
	}
	return nil
}
func (c MongoCollection) Create(ctx context.Context, document interface{}) error {
	client, cll := c.getCollection(ctx)
	defer isError(ctx, client.Disconnect)
	if _, err := cll.InsertOne(ctx, document); err != nil {
		return err
	}
	return nil
}
func (c MongoCollection) Find(ctx context.Context) []interface{} {
	client, cll := c.getCollection(ctx)
	defer isError(ctx, client.Disconnect)
	cursor, err := cll.Find(ctx, bson.M{})
	if err != nil {
		panic(err)
	}
	defer isError(ctx, cursor.Close)
	var packages []interface{}
	for cursor.Next(ctx) {
		var doc interface{}
		if err := cursor.Decode(&doc); err != nil {
			panic(err)
		}
		packages = append(packages, doc)
	}
	return packages
}
