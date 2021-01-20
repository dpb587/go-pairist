package api

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/pkg/errors"
)

type Client struct {
	client    *http.Client
	projectID string
	auth      *Auth
}

func NewClient(client *http.Client, projectID string, auth *Auth) *Client {
	return &Client{
		client:    client,
		projectID: projectID,
		auth:      auth,
	}
}

func (c *Client) get(path string, data interface{}) error {
	uri := fmt.Sprintf("https://us-central1-%s.cloudfunctions.net/%s", c.projectID, path)

	req, err := http.NewRequest("GET", uri, nil)
	if err != nil {
		return err
	}

	if c.auth != nil {
		token, err := c.auth.IDToken()
		if err != nil {
			return err
		}

		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
	}

	res, err := c.client.Do(req)
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

func (c *Client) GetTeamCurrent(team string) (*TeamPairing, error) {
	var res TeamPairing

	err := c.get(fmt.Sprintf("api/current/%s", team), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *Client) GetTeamLists(team string) (*TeamLists, error) {
	var res TeamLists

	err := c.get(fmt.Sprintf("api/lists/%s", team), &res)
	if err != nil {
		return nil, err
	}

	sort.Slice(res.Lists, func(i, j int) bool {
		return res.Lists[i].Order < res.Lists[j].Order
	})

	for _, list := range res.Lists {
		sort.Slice(list.Items, func(i, j int) bool {
			return list.Items[i].Order < list.Items[j].Order
		})
	}

	return &res, nil
}

func (c *Client) GetTeamHistorical(team string) (*TeamPairingHistory, error) {
	var res TeamPairingHistory

	err := c.get(fmt.Sprintf("api/history/%s", team), &res)
	if err != nil {
		return nil, err
	}

	return &res, nil
}
