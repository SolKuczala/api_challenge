package db

import (
	"bet_challenge/pkg/oddsapi"
	"context"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const DATA_BASE = "oddsdb"
const COLLECTION_SPORTS = "sports"
const COLLECTION_ODDS = "odds"

type DBClient struct {
	client *mongo.Client
}

func NewDBClient(conString string) (*DBClient, error) {
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
		log.Error("Failed to list databases")
		return nil, err
	}

	log.WithFields(log.Fields{"dbs": databases}).Info("Found databases")
	return &DBClient{client}, nil
}

func (dbc *DBClient) CreateIndex(collectionName, keyName string) (string, error) {
	log.Info("Creating index: ", DATA_BASE, ":", collectionName, ":", keyName)
	indexView := dbc.client.Database(DATA_BASE).Collection(collectionName).Indexes()

	model := mongo.IndexModel{
		Keys: bson.M{
			"key": 1,
		}, Options: options.Index().SetName("index_" + keyName),
	}

	indexName, err := indexView.CreateOne(context.Background(), model, options.CreateIndexes())
	if err != nil {
		log.Error("Failed to create index")
		return "", err
	}
	return indexName, err
}

func (dbc *DBClient) PrintIndexes(collectionName string) error {
	indexView := dbc.client.Database(DATA_BASE).Collection(collectionName).Indexes()
	cursor, err := indexView.List(context.Background(), options.ListIndexes())
	if err != nil {
		log.Error("Failed to list indexes")
		return err
	}
	var results []bson.M
	if err = cursor.All(context.Background(), &results); err != nil {
		log.Fatal(err)
	}
	for _, result := range results {
		log.Info("Found index: ", result)
	}
	return nil
}

func (dbc *DBClient) SaveSport(sport *oddsapi.Sport) error {
	collection := dbc.client.Database(DATA_BASE).Collection(COLLECTION_SPORTS)
	insertResult, err := collection.InsertOne(context.Background(), sport)
	if err != nil {
		log.Error("Failed to save sport")
		return err
	}
	log.WithFields(log.Fields{"Inserted sport with ID:": insertResult.InsertedID}).Info("Sport saved")
	return nil
}

func (db *DBClient) Close() {
	db.client.Disconnect(context.Background())
}

func (db *DBClient) Update() (string, error) {
	return "new connection", nil
}
