package services

import (
	"covid-us-api/file"
	"covid-us-api/requests"
	"encoding/json"
	"fmt"
	"io/ioutil"
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
	Positive     int `json:"positive"`
	Negative     int `json:"negative"`
	PosNeg       int `json:"posNeg"`
	Hospitalized int `json:"hospitalized"`
	Death        int `json:"death"`
	Total        int `json:"totalTestResults"`
}

type DailyAll struct {
	Daily       []Daily `json:"daily_data"`
	LastUpdated string  `json:"last_updated"`
}

type SummaryAll struct {
	Summary     []Summary `json:"summary"`
	LastUpdated string    `json:"last_updated"`
}

type StateAll struct {
	StateData   map[string][]Daily `json:"state_data"`
	LastUpdated string             `json:"last_updated"`
}
type Covid struct {
	Request *requests.Request
}

func (c *Covid) UploadMainPage() {

	var s3Manageer S3Manager

	s3Manageer.UploadFile("covid-19-us-dataset", "covid.html")
}

func (c *Covid) UploadDatasourcesPage() {
	var s3Manageer S3Manager
	s3Manageer.UploadFile("covid-19-us-dataset", "datasources.html")
}

func (c *Covid) UploadAllStateFiles() {
	files, err := ioutil.ReadDir("./states")
	if err != nil {
		log.Fatal(err)
	}
	for _, f := range files {
		var s3Manageer S3Manager
		fullPath := fmt.Sprintf("%s/%s", "states", f.Name())
		s3Manageer.UploadFile("covid-19-us-dataset", fullPath)
	}
}

func (c *Covid) GenerateNewDailyCasesData() error {
	log.Println("Executing a new api call ..")
	response, err := c.Request.NewGetRequest("/api/states/daily")

	if err != nil {
		return err
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

	err = c.GenerateStateData(dailyValues)
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

func (c *Covid) GenerateStateData(dailyValues []Daily) error {
	stateData := make(map[string][]Daily)

	for _, d := range dailyValues {
		stateData[d.State] = append(stateData[d.State], d)
	}

	loc, err := time.LoadLocation("America/Chicago")
	if err != nil {
		return err

	}

	t := time.Now().In(loc)
	lastUpdated := t.Format(time.RFC822)

	stateAll := StateAll{
		StateData:   stateData,
		LastUpdated: lastUpdated,
	}

	dataToWrite, err := json.Marshal(&stateAll)

	if err != nil {
		return err
	}

	return file.SaveFile("stateData.json", "", dataToWrite)
}
func (c *Covid) GenerateNewOverallCasesData() error {
	log.Println("Executing a new api call ..overall")
	response, err := c.Request.NewGetRequest("/api/us")

	if err != nil {
		log.Println(err)
	}

	fmt.Println(string(response))
	var dailyValues []Summary

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

	d := SummaryAll{
		Summary:     dailyValues,
		LastUpdated: lastUpdated,
	}

	dataToWrite, err := json.Marshal(&d)

	if err != nil {
		return err
	}

	return file.SaveFile("summary.json", "", dataToWrite)
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

func (c *Covid) GetDailyStateDataRefactor() (StateAll, error) {
	readData, err := file.ReadFile("stateData.json", "")
	var stateDailyValues StateAll

	if err != nil {
		log.Printf("Unable to open file : %s", err.Error())
		err = c.GenerateNewDailyCasesData()
		if err != nil {
			return stateDailyValues, err
		}
	} else {
		log.Println("File read successfully no external call needed")
	}

	err = json.Unmarshal(readData, &stateDailyValues)
	if err != nil {
		return stateDailyValues, err
	}

	return stateDailyValues, err
}
func (c *Covid) GetSummaryCasesUSRefactor() (SummaryAll, error) {
	readData, err := file.ReadFile("summary.json", "")

	var overallValues SummaryAll

	if err != nil {
		log.Printf("Unable to open file : %s", err.Error())
		err = c.GenerateNewOverallCasesData()
		if err != nil {
			return overallValues, err
		}
	} else {
		log.Println("File read successfully no external call needed")
	}

	err = json.Unmarshal(readData, &overallValues)
	if err != nil {
		return overallValues, err
	}

	return overallValues, err
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
