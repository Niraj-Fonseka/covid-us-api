package services

import (
	"covid-us-api/file"
	"covid-us-api/requests"
	"encoding/json"
	"fmt"
	"log"
	"time"
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

type Summary struct {
	Positive int `json:"positive"`
	Negative int `json:"negative"`
	PosNeg   int `json:"posNeg"`
	Pending  int `json:"pending"`
	Death    int `json:"death"`
	Total    int `json:"total"`
}

type DailyAll struct {
	Daily       []Daily `json:"daily_data"`
	LastUpdated string  `json:"last_updated"`
}

type Covid struct {
	Request *requests.Request
}

func (c *Covid) GenerateNewDailyCasesData() error {
	log.Println("Executing a new api call ..")
	response, err := c.Request.NewGetRequest("/api/states/daily")

	if err != nil {
		log.Println(err)
	}

	var dailyValues []Daily

	err = json.Unmarshal(response, &dailyValues)

	if err != nil {
		return err
	}

	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return err

	}

	t := time.Now().In(loc)
	lastUpdated := t.Format(time.RFC822)

	d := DailyAll{
		Daily:       dailyValues,
		LastUpdated: lastUpdated,
	}

	dataToWrite, err := json.Marshal(&d)

	if err != nil {
		return err
	}

	return file.SaveFile("daily.json", "", dataToWrite)
}

func (c *Covid) GetDailyCasesUSRefactor() (DailyAll, error) {
	readData, err := file.ReadFile("daily.json", "")
	var dailyValues DailyAll

	if err != nil {
		log.Printf("Unable to open file : %s", err.Error())
		err = c.GenerateNewDailyCasesData()
		if err != nil {
			return dailyValues, err
		}
	} else {
		log.Println("File read successfully no external call needed")
	}

	err = json.Unmarshal(readData, &dailyValues)
	if err != nil {
		return dailyValues, err
	}

	return dailyValues, err
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

func (c *Covid) GetUSSummary() ([]Summary, error) {

	response, err := c.Request.NewGetRequest("/api/us")

	if err != nil {
		return nil, err
	}

	var usSummary []Summary

	err = json.Unmarshal(response, &usSummary)

	if err != nil {
		return nil, err
	}

	return usSummary, nil
}
