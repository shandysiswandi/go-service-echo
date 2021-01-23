package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func mongoConnection() (*mongo.Database, error) {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return nil, err
	}

	return client.Database(os.Getenv("MONGO_DATABASE")), nil
}
