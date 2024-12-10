package database

import (
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"log/slog"
	"os"
)

var clientDatabase *mongo.Client
var AppDatabase *mongo.Database

var UserCollection *mongo.Collection

func ConnectDatabase() {
	clientOptions := options.Client().ApplyURI(os.Getenv("DB_URL"))
	client, err := mongo.Connect(clientOptions)

	if err != nil {
		panic(err)
	}

	slog.Info("Successfully connected to MongoDB")

	AppDatabase = client.Database("app")
	UserCollection = AppDatabase.Collection("users")
}
