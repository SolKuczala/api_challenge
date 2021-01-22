package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

const DEFAULT_BASE_URL = "https://api.the-odds-api.com/v3"

type Client struct {
	apiKey  string
	baseURL string
}

type ResponseWrapper struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

type Sport struct {
	Key          string `json:"key"`
	Active       bool   `json:"active"`
	Group        string `json:"group"`
	Details      string `json:"details"`
	Title        string `json:"title"`
	HasOutrights bool   `json:"has_outrights"`
}

type Odds struct {
	H2H []float64 `json:"h2h"`
}

type Site struct {
	Key        string    `json:"site_key"`
	Nice       string    `json:"site_nice"`
	LastUpdate time.Time `json:"last_update"`
	Odds       Odds      `json:"odds"`
}

type Match struct {
	SportKey     string    `json:"sport_key"`
	SportNice    string    `json:"sport_nice"`
	Teams        []string  `json:"teams"`
	CommenceTime time.Time `json:"commence_time"`
	HomeTeam     string    `json:"home_team"`
	Sites        []Site    `json:"sites"`
	SitesCount   int       `json:"sites_count"`
}

func NewClient(apiKey, baseURL string) (*Client, error) {
	if len(apiKey) < 1 || len(baseURL) < 1 {
		return nil, errors.New("Can't create client. Missin API key or base URL")
	}
	return &Client{apiKey, baseURL}, nil
}

func (c *Client) GetSports() ([]Sport, error) {
	builtURL := fmt.Sprintf("%s/sports/?apiKey=%s", c.baseURL, c.apiKey)
	data, err := c.get(builtURL)
	if err != nil {
		return nil, err
	}
	var sports []Sport
	err = json.Unmarshal(data, &sports)
	if err != nil {
		return nil, err
	}
	return sports, nil
}

//func (c *Client) GetOdds() (*ResponseBody, error) {
//var jsonBodyResp ResponseBody
//builtUrl := fmt.Sprintf("%s/odds/?apiKey=%s&sport=%s&region=%s&mkt=%s", c.baseURL, c.apiKey)
//responseData, _ := c.get(builtUrl)
//err := json.Unmarshal(responseData, &jsonBodyResp)
//if err != nil {
//return nil, err
//}
//return &jsonBodyResp, nil
//}

func (c *Client) get(url string) (json.RawMessage, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response status code != 200. Status:%d", resp.StatusCode)
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respWrapper ResponseWrapper
	err = json.Unmarshal(responseBody, &respWrapper)
	if err != nil {
		return nil, err
	}

	if !respWrapper.Success {
		return nil, errors.New("Response header sucess field is false")
	}

	return respWrapper.Data, nil
}
