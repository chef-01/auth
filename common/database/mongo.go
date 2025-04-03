package database

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoClient *mongo.Client

// InitMongoDB initializes the MongoDB connection.
func InitMongoDB() {
	mongoURI := os.Getenv("MONGO_URI") // MongoDB URI from environment variable.
	if mongoURI == "" {
		log.Fatal("MONGO_URI environment variable is not set")
	}

	// Read the request timeout from environment
	timeoutStr := os.Getenv("REQUEST_TIMEOUT")
	timeoutSec, err := strconv.Atoi(timeoutStr)
	if err != nil {
		timeoutSec = 10
	}
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Use context with timeout for connection.
	ctx, cancel := context.WithTimeout(context.Background(), time.Duration(timeoutSec)*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatalf("Failed to connect to MongoDB: %v", err)
	}

	// Test the connection with a ping.
	pingCtx, pingCancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer pingCancel()
	if err := client.Ping(pingCtx, nil); err != nil {
		log.Fatalf("Failed to ping MongoDB: %v", err)
	}

	MongoClient = client
	fmt.Println("MongoDB connection established")
}