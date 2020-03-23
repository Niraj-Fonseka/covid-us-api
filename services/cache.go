package services

import "fmt"

type Cache struct {
	CovidService *Covid
}

func NewCache(covidService *Covid) *Cache {

	fmt.Println("Creating new cache : ", covidService)
	return &Cache{
		CovidService: covidService,
	}
}

func (c *Cache) GetDailyRecordsByDate(date int) (DailyAll, error) {

	fmt.Println("InGetDailyRecordsByDate  : covidService", c)
	return c.CovidService.GetDailyCasesUSRefactor()

}

func (c *Cache) GetDailyStateRecords() (StateAll, error) {
	return c.CovidService.GetDailyStateDataRefactor()
}

func (c *Cache) CreateOverallRecords(filename string) {

}

func (c *Cache) CreateDailyRecords(filename string) {

}

func (c *Cache) GetOverallRecords() (SummaryAll, error) {
	return c.CovidService.GetSummaryCasesUSRefactor()
}
