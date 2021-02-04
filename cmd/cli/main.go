package main

import (
	"bet_challenge/internal/db"
	"bet_challenge/pkg/oddsapi"
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"
)

func main() {
	apiKey, dbConString, minutes := getEnv()

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

	odds, err := client.GetOdds("upcoming", "uk", "h2h")
	if err != nil {
		log.Error(err)
	}

	for _, odd := range odds {
		DB.SaveMatch(&odd)
	}

	DB.UpdateOddsEvery(minutes, odds)
}

func getEnv() (string, string, int) {
	apiKey := os.Getenv("API_KEY")
	dbConString := os.Getenv("DB_CON_STRING")
	time := os.Getenv("MINUTES")
	if len(apiKey) < 1 || len(apiKey) < 1 || len(time) < 1 {
		log.Fatal("Invalid env variables")
	}
	minutes, err := strconv.Atoi(time)
	if err != nil {
		log.Fatal("Could not get time for update", err)
	}
	return apiKey, dbConString, minutes
}
