package db

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DB struct {
	DB   string
	Conn *mongo.Client
}

func StartDatabase(URI, Database string) *DB {

	c, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(URI))
	if err != nil {
		return &DB{}
	}

	err = c.Ping(context.TODO(), nil)
	if err != nil {
		return &DB{}
	}

	return &DB{
		DB:   Database,
		Conn: c,
	}
}
