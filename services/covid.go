package services

import (
	"covid-us-api/requests"
	"encoding/json"
	"fmt"
)

type Daily struct {
	Date     int    `json:"date"`
	State    string `json:"state"`
	Positive int    `json:"positive"`
	Negative int    `json:"negative"`
	Pending  int    `json:"pending"`
	Death    int    `json:"death"`
	Total    int    `json:"total"`
}

type Covid struct {
	Request *requests.Request
}

func (c *Covid) GetDailyCasesUS() ([]Daily, error) {
	response, err := c.Request.NewGetRequest("/api/states/daily")

	if err != nil {
		return nil, err
	}

	var dailyValues []Daily

	err = json.Unmarshal(response, &dailyValues)

	if err != nil {
		return nil, err
	}

	return dailyValues, nil
}

func (c *Covid) GetDailyCasesUSByState(state string) ([]Daily, error) {

	buildURL := fmt.Sprintf("/api/states/daily?state=%s", state)

	response, err := c.Request.NewGetRequest(buildURL)

	if err != nil {
		return nil, err
	}

	var dailyValues []Daily

	err = json.Unmarshal(response, &dailyValues)

	if err != nil {
		return nil, err
	}

	return dailyValues, nil

}
