package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// DB Core structure for database user
type DB struct {
	Client   *mongo.Client
	URI      string
	Database string
}

// StartDatabase Will start a database instance with the specified uri and params
func StartDatabase(database string) *DB {

	uri := os.Getenv("DB")
	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return nil
	}

	return &DB{
		URI:      uri,
		Client:   c,
		Database: database,
	}
}
