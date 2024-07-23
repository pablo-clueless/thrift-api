package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() mongo.Client {
	uri := "mongodb+srv://smsnmicheal:pxAGTm9tbpMJSfzG@cluster0.spekxgn.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		panic(err)
	}
	fmt.Println("connection to database established")

	return *client
}

var Database mongo.Client = Connect()

func GetCollection(client mongo.Client, collectionName string) *mongo.Collection {
	return client.Database("thrift").Collection(collectionName)
}
