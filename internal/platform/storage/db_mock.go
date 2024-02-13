package storage

import (
	"context"

	"github.com/agustinrabini/go-toolkit/gonosql"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func NewMock() *gonosql.Data {
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017").SetDirect(true))
	database := client.Database("testdb")
	return &gonosql.Data{
		DB:       client,
		Database: database,
		Error:    err,
	}
}
