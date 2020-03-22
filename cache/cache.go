package cache

import (
	"covid-us-api/services"
)

type Cache struct {
	Daily   map[int]map[string]services.Daily
	Overall map[string]int
}

func (c *Cache) GetDailyRecordsByDate(date int) {

}

func (c *Cache) GetDailyRecordsByState(state string) {

}

func (c *Cache) CreateOverallRecords(filename string) {

}

func (c *Cache) CreateDailyRecords(filename string) {

}

func (c *Cache) GetOverallRecords() services.Summary {

	return services.Summary{}
}
