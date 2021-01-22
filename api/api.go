package api

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
)

const DEFAULT_BASE_URL = "https://api.the-odds-api.com/v3/"

//puede un oddsapiclient tener directamente como metodo el get?
type OddsAPIClient struct {
	apiKey  string
	baseURL string
}

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

func NewOddsAPIClient(apiKey, baseURL string) (*OddsAPIClient, error) {
	if len(apiKey) < 1 || len(baseURL) < 0 {
		return nil, errors.New("Can't create client. Missin API key or base URL")
	}

	return &OddsAPIClient{
		apiKey: apiKey, baseURL: baseURL,
	}, nil
}

//tiene que recibir solo los parametros, construye solo el path
func GetSports(params OddsAPIClient) (*ResponseBody, error) {
	var jsonBodyResp ResponseBody
	builtUrl := params.baseURL + "sports/?apiKey=" + params.apiKey
	responseData, _ := Get(builtUrl)
	err := json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		return nil, err
	}
	return &jsonBodyResp, nil
}

func GetOdds(params OddsAPIClient) (*ResponseBody, error) {
	var jsonBodyResp ResponseBody
	builtUrl := params.baseURL + "odds/?apiKey=" + params.apiKey + "&sport={sport}&region={region}&mkt={mkt}"
	responseData, _ := Get(builtUrl)
	err := json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		return nil, err
	}
	return &jsonBodyResp, nil
}

func Get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
