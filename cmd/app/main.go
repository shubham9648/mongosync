package main

import (
	"github.com/shubham9648/mongosync/internal/mongo"
	"go.uber.org/zap"
)

func main() {

	client, err := mongo.ConnectToMongo()

	if err != nil {
		zap.L().Error("error while connecting to mongoDB", zap.Error(err))
		panic(err)
	}
}
