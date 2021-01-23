package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var mongoCon *mongo.Database

func mongoConnection() error {
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URI"))
	client, err := mongo.NewClient(clientOptions)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		return err
	}

	mongoCon = client.Database(os.Getenv("MONGO_DATABASE"))

	return nil
}

// GetMongoDB is
func GetMongoDB() *mongo.Database {
	return mongoCon
}
