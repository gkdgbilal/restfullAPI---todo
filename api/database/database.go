package database

import (
	"RestFullAPI-todo/configs"
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Connect() (*mongo.Database, error) {
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	client, err := mongo.Connect(ctx, options.Client().
		ApplyURI(fmt.Sprintf(configs.DBConnectionURL())))
	if err != nil {
		return nil, err
	}
	err = client.Ping(context.Background(), nil)
	if err != nil {
		return nil, err
	}
	db := client.Database(configs.C.Database.Name)

	return db, nil
}

var (
	Users = "users"
	Todos = "todos"
)
