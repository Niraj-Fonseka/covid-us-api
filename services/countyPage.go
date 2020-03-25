package services

import (
	"encoding/json"
	"fmt"
	"strings"
)

type CountyPage struct {
	CovidService *Covid
	CacheService *Cache
}

type CountyRecord struct {
	Code  string `json"code"`
	Name  string `json"name"`
	Value int    `json"value"`
}

func (c *CountyPage) GenerateData(countyData USCountyAll) (map[string]interface{}, error) {

	dataStore := make(map[string]interface{})
	var generatedDataConfirmed []CountyRecord
	var generatedDataDeath []CountyRecord

	for state, countyList := range countyData.CountyData {
		stateLower := strings.ToLower(state)
		for _, county := range countyList {
			code := generateStateCode(stateLower, county.County)
			generatedName := generateCountyName(county.County, state)

			generatedDataDeath = append(generatedDataDeath, CountyRecord{Code: code, Name: generatedName, Value: county.Deaths[len(county.Deaths)-1]})
			generatedDataConfirmed = append(generatedDataConfirmed, CountyRecord{Code: code, Name: generatedName, Value: county.Confirmed[len(county.Confirmed)-1]})

		}
	}

	confirmedJSON, err := json.Marshal(generatedDataConfirmed)

	if err != nil {
		return dataStore, err
	}

	deathsJSON, err := json.Marshal(generatedDataDeath)
	if err != nil {
		return dataStore, err
	}
	dataStore["deathsData"] = deathsJSON
	dataStore["confirmedData"] = confirmedJSON

}

func generateStateCode(state, code string) string {
	state = strings.ToLower(state)
	shortenedCode := code[2:]

	return fmt.Sprintf("us-%s-%s", state, shortenedCode)
}

func generateCountyName(county, state string) string {
	return fmt.Sprintf("%s , %s", strings.TrimSpace(county), strings.TrimSpace(state))
}
