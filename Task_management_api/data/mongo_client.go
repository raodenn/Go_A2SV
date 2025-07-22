package data

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var TaskCollections *mongo.Collection
var MongoClient *mongo.Client
var UserCollection *mongo.Collection

func ConnectToMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		panic(fmt.Sprintf("MongoDB connection failed: %v", err))
	}

	if err := client.Ping(ctx, nil); err != nil {
		panic(fmt.Sprintf("MongoDB ping failed: %v", err))
	}

	MongoClient = client
	TaskCollections = client.Database("taskdb").Collection("tasks")
	UserCollection = client.Database("taskdb").Collection("users")

	fmt.Println("Connected to MongoDB and selected collection")
}
