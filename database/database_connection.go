package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
)

// Mogodb collection = table  flow Project -> DB -> collection
func DBinstance() *mongo.Client {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error Loading .env file")
	}
	mongoDb := os.Getenv("MONGODB_URL")

	client, err := mongo.Connect(options.Client().ApplyURI(mongoDb).SetConnectTimeout(10 * time.Second))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Connect to Mongo DB")
	return client
}

var Client *mongo.Client = DBinstance()

func OpenCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	var collection *mongo.Collection = client.Database("JWT").Collection(collectionName)
	return collection
}

//	defer func() {
//		if err := client.Disconnect(context.TODO()); err != nil {
//			panic(err)
//		}
//	}()
func CloseDB() {
	if err := Client.Disconnect(context.TODO()); err != nil {
		log.Fatal("Error closing MongoDB connection:", err)
	}
	fmt.Println("MongoDB connection closed")
}
