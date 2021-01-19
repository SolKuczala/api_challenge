package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type ResponseBody struct {
	Success bool `json:"success"`
	Data    []struct {
		Key          string `json:"key"`
		Active       bool   `json:"active"`
		Group        string `json:"group"`
		Details      string `json:"details"`
		Title        string `json:"title"`
		HasOutrights bool   `json:"has_outrights"`
	} `json:"data"`
}

var apiKEY = "c451435480578edc86c4de5e3d6275cf" //variable de entorno
//var dataBaseConnection = "string"               //no se bien como, variable de entorno

func main() {
	var jsonBodyResp ResponseBody
	//keywords = "developer"
	//api := os.Getenv(apiKEY)
	url := "https://api.the-odds-api.com/v3/sports/?apiKey=" + apiKEY
	fmt.Println("good") //https://api.the-odds-api.com/v3/sports/?apiKey=c451435480578edc86c4de5e3d6275cf
	resp, err := http.Get(url)
	if err != nil {
		println(err)
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err.Error())
		os.Exit(1)
	}

	err = json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	fmt.Println(jsonBodyResp.Data)
}
