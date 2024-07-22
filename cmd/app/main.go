package main

import (
	"fmt"

	es "github.com/shubham9648/mongosync/internal/elasticsearch"
	"github.com/shubham9648/mongosync/internal/mongo"
	"go.uber.org/zap"
)

func main() {

	client, err := mongo.ConnectToMongo()

	if err != nil {
		zap.L().Error("error while connecting to mongoDB", zap.Error(err))
		panic(err)
	}
	fmt.Println("client is ", client)

	esClient, err := es.InitializeESClient()
	fmt.Println("client is ", esClient)
}
