package driver

import (
	"base-app/config"
	"base-app/utils"
	"context"
	"log"
	"sync"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.uber.org/zap"
)

var once sync.Once
var dbClient *mongo.Client

// Mongo db driver to manage mongo db connection
func InitializeDbConnection() *mongo.Client {
	Logger := utils.Logger

	once.Do(func() {
		connCtx, _ := context.WithTimeout(context.Background(), 5*time.Second)
		client, err := mongo.Connect(connCtx, options.Client().ApplyURI(config.AppConfig.Database.Uri))
		if err != nil {
			log.Fatal("Failed to establish DB connection", err)
			return
		}

		Logger.Info("Database client created")
		dbClient = client
		if ok, conErr := connectionHealthCheck(); !ok {
			log.Fatal("DB connection error ", conErr)
		}
	})

	return dbClient

}

// function to fetch dbConnection
func GetDBClient() (dbClient *mongo.Client) {

	connectionHealthCheck()
	return dbClient

}

func connectionHealthCheck() (bool, error) {
	Logger := utils.Logger
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	pingErr := dbClient.Ping(ctx, nil)
	if pingErr != nil {
		Logger.Error("Db Connection error: ", zap.Any("Error", pingErr))
		return false, pingErr
	}
	Logger.Info("DB Connection is healthy")
	return true, nil
}
