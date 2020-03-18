package requests

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

type Request struct {
	Client  *http.Client
	BaseURL string
}

func (r *Request) NewGetRequest(route string) ([]byte, error) {

	builtURL, _ := buildURL(r.BaseURL, route)

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
