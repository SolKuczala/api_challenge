package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

//var dataBaseConnection = "string"//no se bien como, variable de entorno

func main() {

	apiKEY := os.Getenv("API_KEY") //variable de entorno
	println("pase por getenv")

	url := "https://api.the-odds-api.com/v3/sports/?apiKey=" + apiKEY
	jsonBodyResp, err := GetSports(url)
	if err != nil {
		println(err)
	}
	fmt.Println(jsonBodyResp)
}

func GetSports(url string) (*ResponseBody, error) {
	var jsonBodyResp ResponseBody
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err

	}

	err = json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		return nil, err

	}
	return &jsonBodyResp, nil
}
