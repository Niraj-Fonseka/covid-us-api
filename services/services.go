package services

import (
	"covid-us-api/requests"
	"net/http"
)

type Services struct {
	Covid *Covid
	Graph *Graph
}

func RegisterServices() *Services {
	request := requests.Request{
		Client:  &http.Client{},
		BaseURL: "https://covidtracking.com",
	}

	return &Services{
		Covid: &Covid{
			Request: &request,
		},
		Graph: &Graph{},
	}
}
