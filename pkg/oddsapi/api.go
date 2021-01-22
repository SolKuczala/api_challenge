package oddsapi

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/sirupsen/logrus"
)

type Client struct {
	apiKey  string
	baseURL string
}

func NewClient(apiKey, baseURL string) (*Client, error) {
	if len(apiKey) < 1 || len(baseURL) < 1 {
		return nil, errors.New("Can't create client. Missin API key or base URL")
	}
	return &Client{apiKey, baseURL}, nil
}

func (c *Client) GetSports() ([]Sport, error) {
	return c.GetSportsCustom(DEFAULT_ALL_VALUE, DEFAULT_OUTRIGHTS_VALUE)
}

func (c *Client) GetSportsCustom(all, outrights bool) ([]Sport, error) {
	builtURL := fmt.Sprintf(
		"%s/sports/?apiKey=%s&all=%t&outrights=%t",
		c.baseURL,
		c.apiKey,
		all,
		outrights,
	)
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

func (c *Client) GetOdds(sport, region, market string) ([]Match, error) {
	return c.GetOddsCustom(sport, region, market, DEFAULT_DATE_FORMAT, DEFAULT_ODDS_FORMAT)
}

func (c *Client) GetOddsCustom(sport, region, market, dateFormat, oddsFormat string) ([]Match, error) {
	builtURL := fmt.Sprintf(
		"%s/odds/?apiKey=%s&sport=%s&region=%s&mkt=%s&dateFormat=%s&oddsFormat=%s",
		c.baseURL,
		c.apiKey,
		sport,
		region,
		market,
		dateFormat,
		oddsFormat,
	)
	data, err := c.get(builtURL)
	if err != nil {
		return nil, err
	}
	var matches []Match
	err = json.Unmarshal(data, &matches)
	if err != nil {
		return nil, err
	}
	return matches, nil
}

type responseWrapper struct {
	Success bool            `json:"success"`
	Data    json.RawMessage `json:"data"`
}

func (c *Client) get(url string) (json.RawMessage, error) {
	log.Info("Calling: ", url)
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.WithFields(log.Fields{"status_code": resp.StatusCode}).Error("Request failed")
		return nil, ERRRequestStatusCodeNotOk
	}

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var respWrapper responseWrapper
	err = json.Unmarshal(responseBody, &respWrapper)
	if err != nil {
		return nil, err
	}

	if !respWrapper.Success {
		return nil, ERRSuccessFieldIsFalse
	}

	return respWrapper.Data, nil
}
