package main

import (
	"bet_challenge/internal/db"
	"bet_challenge/pkg/oddsapi"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	apiKey, dbConString := getEnv()

	DB, err := db.NewDBClient(dbConString)
	if err != nil {
		log.Fatal(err)
	}
	defer DB.Close()

	client, err := oddsapi.NewClient(apiKey, oddsapi.DEFAULT_BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	sports, err := client.GetSports()
	if err != nil {
		log.Error(err)
	}

	for _, sport := range sports {
		DB.SaveSport(&sport)
	}
	//defineDBIndexes(DB)
	odds, err := client.GetOdds("upcoming", "uk", "h2h")
	if err != nil {
		log.Error(err)
	}

	for _, odd := range odds {
		DB.SaveMatch(&odd)
	}

	DB.UpdateOddsEvery(60, odds)
}

func getEnv() (string, string) {
	apiKey := os.Getenv("API_KEY")
	dbConString := os.Getenv("DB_CON_STRING")
	if len(apiKey) < 1 || len(apiKey) < 1 {
		log.Fatal("Invalid env variables")
	}
	return apiKey, dbConString
}

//func defineDBIndexes(DB *db.DBClient) {
//	for collectionName, keys := range map[string]bson.M{
//		db.COLLECTION_ODDS: bson.M{"sport_key": 1, "home_team": 1, "commence_time": 1},
//	} {
//		indexName, err := DB.CreateIndex(collectionName, keys)
//		if err != nil {
//			log.Error("Failed to create index: ", collectionName, ":", keys)
//			log.Error(err)
//		}
//		log.Info("Created index: ", indexName)
//		err = DB.PrintIndexes(collectionName)
//		if err != nil {
//			log.Error(err)
//		}
//	}
//}
