package common

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func NewMongoConnection(ctx context.Context) *mongo.Client {
	opts := options.Client().ApplyURI(MongoUri)

	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		panic(err)
	}

	fmt.Println("Mongo connection acquired")

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		panic(err)
	}

	fmt.Println("Mongo primary server discovered")

	return client
}
