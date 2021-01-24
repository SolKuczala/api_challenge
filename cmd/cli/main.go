package main

import (
	"bet_challenge/internal/db"
	"bet_challenge/pkg/oddsapi"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {
	apiKey := os.Getenv("API_KEY")
	dbConString := os.Getenv("DB_CON_STRING")
	if len(apiKey) < 1 || len(apiKey) < 1 {
		log.Fatal("Invalid env variables")
	}

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

	odds, err := client.GetOdds("soccer_epl", "uk", "h2h")
	if err != nil {
		log.Error(err)
	}
	for _, odds := range odds {
		log.Info(odds)
	}

	for collectionName, keyName := range map[string]string{
		db.COLLECTION_SPORTS: "key",
		db.COLLECTION_ODDS:   "sport_key",
	} {
		indexName, err := DB.CreateIndex(collectionName, keyName)
		if err != nil {
			log.Error("Failed to create index: ", collectionName, ":", keyName)
			log.Error(err)
		}
		log.Info("Created index: ", indexName)
		err = DB.PrintIndexes(collectionName)
		if err != nil {
			log.Error(err)
		}
	}

}
