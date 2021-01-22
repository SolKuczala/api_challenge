package main

import (
	"bet_challenge/api"
	"bet_challenge/db"
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

	client, err := api.NewClient(apiKey, api.DEFAULT_BASE_URL)
	if err != nil {
		log.Fatal(err)
	}

	sports, err := client.GetSports()
	if err != nil {
		log.Error(err)
	}
	for _, sport := range sports {
		log.Info(sport)
	}
}
