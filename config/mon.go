package config

import (
	"context"
	"log"
	"os"
	"sync"

	_ "github.com/joho/godotenv/autoload"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
    clientInstance *mongo.Client
    once           sync.Once
)

func GetMongoDatabase() *mongo.Database {
    once.Do(func() {
        uri := os.Getenv("DUNGEON_MANAGER_CONNECTION_STRING")
        if uri == "" {
            log.Fatal("Set your 'DUNGEON_MANAGER_CONNECTION_STRING' environment variable. " +
                "See: " +
                "www.mongodb.com/docs/drivers/go/current/usage-examples/#environment-variable")
        }
        client, err := mongo.Connect(context.TODO(), options.Client().
            ApplyURI(uri))
        if err != nil {
            panic(err)
        }
        clientInstance = client
    })
    return clientInstance.Database("dungeonManagerDB")
}