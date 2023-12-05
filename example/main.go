package example

import (
	"context"
	"github.com/itzngga/mongolog"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

func main() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017/")
	clientOptions.SetMonitor(mongolog.New()) // <- the logger

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	client.Database("test")
}
