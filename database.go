package main

import (
	"context"
	"time"

	"github.com/projectdde/issue-bot/config"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type DataBase struct {
	client *mongo.Client
	ctx context.Context
	settings *config.Yaml
}

func (db *DataBase)open() error {
	opts := options.Client().ApplyURI(*db.settings.Database.URL)
	opts.Auth = &options.Credential{
			Username: *db.settings.Database.Auth.Name,
			Password: *db.settings.Database.Auth.Password,
	}
	client, err := mongo.NewClient(opts)
	if err != nil {
		return err
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)
	db.client = client
	db.ctx = ctx
	client.Connect(ctx)
	return nil
}

func (db *DataBase)close() {
	db.client.Disconnect(db.ctx)
}

func (db *DataBase)Ping() {
	err := db.open()
	if err != nil {
		logrus.Error(err)
		return
	}
	defer db.close()
	ar, _ := db.client.ListDatabaseNames(db.ctx, bson.D{})
	logrus.Infof("Database: %v", ar)
}

// Save data
func (db *DataBase) Save(data interface{}) (*mongo.InsertOneResult, error) {
	db.open()
	defer db.close()
	collection := db.client.Database("Issue").Collection("Issue")
	return collection.InsertOne(db.ctx, data)
}
