package database

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

var Db *mongo.Collection

func Dbconnect() (*mongo.Collection, error) {
	serverApi := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("connection-key").SetServerAPIOptions(serverApi)

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	client, err := mongo.Connect(ctx, opts)
	if err != nil {
		return nil, err
	}

	quickstartDatabase := client.Database("LogApi00")
	logCollections := quickstartDatabase.Collection("logApiv0")
	if err := client.Database("LogApi00").RunCommand(context.TODO(), bson.D{{"ping", 1}}).Err(); err != nil {
		return nil, fmt.Errorf("ping failed: %v", err)
	}
	fmt.Println("Pinged your deployment. You successfully connected to MongoDB!")

	return logCollections, nil
}
