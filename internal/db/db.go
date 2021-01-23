package db

import (
	"context"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type DBClient struct {
	client *mongo.Client
}

func NewDBClient(conString string) (*DBClient, error) {
	log.Infof("Connecting to %s", conString)
	client, err := mongo.NewClient(options.Client().ApplyURI(conString))
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	err = client.Connect(ctx)
	if err != nil {
		log.Error("Failed to connect to DB")
		return nil, err
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Error("Failed to ping DB")
		return nil, err
	}

	databases, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		log.Error("Failed to list databses")
		return nil, err
	}
	log.WithFields(log.Fields{"dbs": databases}).Info("Found databses")

	return &DBClient{client}, nil
}

func (db *DBClient) Save() (string, error) {
	return "new connection", nil
}

func (db *DBClient) Close() {
	return
}

func (db *DBClient) Update() (string, error) {
	return "new connection", nil
}
