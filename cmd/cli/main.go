package main

import (
	"bet_challenge/internal/db"
	"bet_challenge/pkg/oddsapi"
	"os"
	"time"

	log "github.com/sirupsen/logrus"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	apiKey, dbConString := getEnv()

	DB, err := db.NewDBClient(dbConString)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()
	defineDBIndexes(DB)

	client, err := oddsapi.NewClient(apiKey, oddsapi.DEFAULT_BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	////test
	sports := []oddsapi.Sport{
		oddsapi.Sport{
			Key:    "americanfootball_nfl",
			Active: true,
			Group:  "American Football"},
	}
	//guardo
	for _, sport := range sports {
		DB.SaveSport(&sport)
	}

	sports = []oddsapi.Sport{
		oddsapi.Sport{
			Key:    "americanfootball_nfl",
			Active: true,
			Group:  "susiQuiu"},
		oddsapi.Sport{
			Key:    "soccer_international",
			Active: true,
			Group:  "Random futbol"},
	}
	for _, sport := range sports {
		DB.SaveSport(&sport)
	}
	////test
	os.Exit(0)

	sports, err = client.GetSports()
	if err != nil {
		log.Error(err)
	}

	for _, sport := range sports {
		DB.SaveSport(&sport)
		log.Info(sport)
	}

	odds, err := client.GetOdds("soccer_epl", "uk", "h2h")
	if err != nil {
		log.Error(err)
	}

	for _, odd := range odds {
		DB.SaveOdds(&odd)
		log.Info(odd)
	}
	for range time.Tick(time.Minute * 1) {
		//matches which are not in-play should be updated every hour
	}
}

func getEnv() (string, string) {
	apiKey := os.Getenv("API_KEY")
	dbConString := os.Getenv("DB_CON_STRING")
	if len(apiKey) < 1 || len(apiKey) < 1 {
		log.Fatal("Invalid env variables")
	}
	return apiKey, dbConString
}

func defineDBIndexes(DB *db.DBClient) {
	for collectionName, keys := range map[string]bson.M{
		db.COLLECTION_SPORTS: bson.M{"key": 1},
		db.COLLECTION_ODDS:   bson.M{"sport_key": 1, "teams": 1, "commence_time": 1},
	} {
		indexName, err := DB.CreateIndex(collectionName, keys)
		if err != nil {
			log.Error("Failed to create index: ", collectionName, ":", keys)
			log.Error(err)
		}
		log.Info("Created index: ", indexName)
		err = DB.PrintIndexes(collectionName)
		if err != nil {
			log.Error(err)
		}
	}
}
