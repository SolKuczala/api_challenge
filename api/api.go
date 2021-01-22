package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
)

const DEFAULT_BASE_URL = "https://api.the-odds-api.com/v3"

type Client struct {
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

func NewClient(apiKey, baseURL string) (*Client, error) {
	if len(apiKey) < 1 || len(baseURL) < 1 {
		return nil, errors.New("Can't create client. Missin API key or base URL")
	}
	return &Client{apiKey, baseURL}, nil
}

func (c *Client) GetSports() (*ResponseBody, error) {
	var jsonBodyResp ResponseBody
	builtUrl := fmt.Sprintf("%s/sports/?apiKey=%s", c.baseURL, c.apiKey)
	responseData, err := c.get(builtUrl)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		return nil, err
	}
	return &jsonBodyResp, nil
}

func (c *Client) GetOdds() (*ResponseBody, error) {
	var jsonBodyResp ResponseBody
	builtUrl := fmt.Sprintf("%s/odds/?apiKey=%s&sport=%s&region=%s&mkt=%s", c.baseURL, c.apiKey)
	responseData, _ := c.get(builtUrl)
	err := json.Unmarshal(responseData, &jsonBodyResp)
	if err != nil {
		return nil, err
	}
	return &jsonBodyResp, nil
}

func (c *Client) get(url string) ([]byte, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("Response status code != 200. Status:%d", resp.StatusCode)
	}

	responseData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	return responseData, nil
}
