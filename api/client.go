package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/pkg/errors"
)

var DefaultClient = NewClient(http.DefaultClient, DefaultDatabaseURL, nil)

type Client struct {
	client  *http.Client
	baseURL string
	auth    *Auth
}

func NewClient(client *http.Client, baseURL string, auth *Auth) *Client {
	return &Client{
		client:  client,
		baseURL: baseURL,
		auth:    auth,
	}
}

func (c *Client) get(path string, data interface{}) error {
	uri := fmt.Sprintf("%s/%s", c.baseURL, path)
	if c.auth != nil {
		token, err := c.auth.IDToken()
		if err != nil {
			return err
		}

		uri = fmt.Sprintf("%s?auth=%s", uri, url.QueryEscape(token))
	}

	res, err := c.client.Get(uri)
	if err != nil {
		return err
	}

	bytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return errors.Wrap(err, "reading response body")
	}

	var errcheck struct {
		Error string `json:"error"`
	}

	// ignore error since successful response data may not be unmarshallable
	_ = json.Unmarshal(bytes, &errcheck)
	if errcheck.Error != "" {
		return errors.Wrap(errors.New(errcheck.Error), "verifying response")
	}

	err = json.Unmarshal(bytes, data)
	if err != nil {
		return errors.Wrap(err, "unmarshalling response")
	}

	return nil
}

func (c *Client) GetTeamCurrent(team string) (*TeamHistorical, error) {
	var res TeamHistorical

	err := c.get(fmt.Sprintf("teams/%s/current.json", team), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamLists(team string) (*TeamLists, error) {
	var res TeamLists

	err := c.get(fmt.Sprintf("teams/%s/lists.json", team), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

type TeamHistoricalFull map[int]TeamHistorical

func (c *Client) GetTeamHistorical(team string) (*TeamHistoricalFull, error) {
	var res TeamHistoricalFull

	err := c.get(fmt.Sprintf("teams/%s/history.json", team), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
