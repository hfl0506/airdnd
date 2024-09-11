package db

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitDB() *mongo.Database {
	uri := os.Getenv("MONGODB_URI")
	dbName := os.Getenv("MONGODB_NAME")

	if uri == "" {
		log.Fatal("Set your 'MONGODB_URI' environment variable")
	}

	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))

	if err != nil {
		panic(err)
	}

	// defer func() {
	// 	if err := client.Disconnect(context.TODO()); err != nil {
	// 		panic(err)
	// 	}
	// }()
	
	db := client.Database(dbName)

	return db
}