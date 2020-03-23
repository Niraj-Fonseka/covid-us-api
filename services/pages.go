package services

import (
	"log"
)

type Pages struct {
	Pages        map[string]PageFramework
	CovidService *Covid
	CacheService *Cache
}

func NewPages(svcs *Services) *Pages {

	p := Pages{
		Pages:        make(map[string]PageFramework),
		CovidService: svcs.Covid,
		CacheService: svcs.Cache,
	}

	p.RegisterPages()
	return &p
}

//Manually register each page
func (p *Pages) RegisterPages() {
	log.Println("Registering Pages..")
	covidPage := CovidPage{p.CovidService, p.CacheService}
	p.RegisterNewPage("newcovid", &covidPage)

	statePage := StatePage{p.CovidService, p.CacheService}
	p.RegisterNewPage("statepage", &statePage)
}
func (p *Pages) RegisterNewPage(pageName string, page PageFramework) {
	p.Pages[pageName] = page
}

func (p *Pages) RenderPages(pageName string) {
	if pageName == "" {

		//render everything
		for _, p := range p.Pages {
			err := p.BuildPage()
			if err != nil {
				log.Printf("Unable to render page for page : %s , error : %s", pageName, err.Error())
			}
		}
	}
}
