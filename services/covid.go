package services

import (
	"covid-us-api/requests"
	"encoding/json"
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
