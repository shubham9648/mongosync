package mongo

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

func ConnectToMongo() (*mongo.Client, error) {

	clientOptions := options.Client().ApplyURI(os.Getenv("MONGO_URL"))
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		zap.L().Error("Failed to create new MongoDB client", zap.Error(err))
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)

	if err != nil {
		zap.L().Error("Failed to connect to MongoDB", zap.Error(err))
		return nil, err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		zap.L().Error("Failed to ping MongoDB", zap.Error(err))
		return nil, err
	}

	zap.L().Info("Connected to MongoDB")
	return client, nil

}
