package requests

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Request struct {
	Client               *http.Client
	BaseURLCovidTracking string
	CountyTrackingURL    string
}

func (r *Request) NewGetRequest(url, route string) ([]byte, error) {

	builtURL, _ := buildURL(url, route)

	//TO DO : Error check

	request, err := http.NewRequest("GET", builtURL, nil)
	if err != nil {
		return nil, err
	}

	request.Header.Set("Content-type", "application/json")

	if err != nil {
		return nil, err
	}

	resp, err := r.Client.Do(request)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		log.Printf("%s : %d \n", route, resp.StatusCode)
		return nil, fmt.Errorf("non 200 response code : %d for route : %s ", resp.StatusCode, route)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}

func buildURL(base string, route string) (string, error) {
	return fmt.Sprintf(base + route), nil //need to do build error checking with slashes and stuff
}
