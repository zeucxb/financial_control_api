package repository

import (
	"context"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client
var database *mongo.Database

// Init initialize repository connecting to database
func Init() {
	ctx, cancelCtx := context.WithTimeout(context.Background(), 60*time.Second)
	defer cancelCtx()

	var err error

	client, err = mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv("MONGO_ATLAS_CLUSTER")))

	if err != nil {
		log.Fatal("Couldn't connect to database - ", err)
		return
	}

	database = client.Database("financial_control")
}
