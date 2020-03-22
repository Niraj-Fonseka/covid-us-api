package services

import (
	"covid-us-api/requests"
	"net/http"
)

type Services struct {
	Covid *Covid
	Graph *Graph
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
	graphService := &Graph{}
	cache := NewCache(covidService)

	return &Services{
		Covid: covidService,
		Graph: graphService,
		Cache: cache,
	}
}
