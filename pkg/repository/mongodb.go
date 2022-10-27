package repository

import (
	"context"
	"os"
	"time"

	notemaker "github.com/abdrakhmanovzh/notemaker2.0"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoInstance struct {
	Client *mongo.Client
	Db     *mongo.Database
}

func Connect() (*mongo.Database, error) {
	notemaker.LoadEnvVariables()

	client, err := mongo.NewClient(options.Client().ApplyURI(os.Getenv("MONGO_URI")))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	db := client.Database(os.Getenv("DB"))

	if err != nil {
		panic(err)
	}

	return db, err
}
