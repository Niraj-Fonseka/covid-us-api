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
		Client:               &http.Client{},
		BaseURLCovidTracking: "https://covidtracking.com",
		CountyTrackingURL:    "https://usafactsstatic.blob.core.windows.net/public/2020/coronavirus-timeline/allData.json",
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
