package services

import (
	"covid-us-api/requests"
	"net/http"
)

type Services struct {
	Covid *Covid
	Cache *Cache
}

func RegisterServices() *Services {
	request := requests.Request{
		Client:  &http.Client{},
		BaseURL: "https://covidtracking.com",
	}

	covidService := &Covid{
		Request: &request,
	}
	cache := NewCache(covidService)

	return &Services{
		Covid: covidService,
		Cache: cache,
	}
}
