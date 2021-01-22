package main

import (
	"bet_challenge/api"
	"os"
)

//var dataBaseConnection = "string"//no se bien como, variable de entorno

func main() {
	apiKey := os.Getenv("API_KEY") //variable de entorno
	db.NewDBClient()
	defer db.Close()
	client, _ := api.NewOddsAPIClient(apiKey, api.DEFAULT_BASE_URL)
	sports, _ := api.GetSports(*client)
	db.Store(sports)

	//odds, err := api.GetOdds(*client)

}
