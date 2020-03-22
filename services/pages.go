package services

import (
	"fmt"
	"log"
)

type Pages struct {
	Pages        map[string]PageFramework
	CovidService *Covid
	CacheService *Cache
}

func NewPages(covidService *Covid) *Pages {
	fmt.Println("creating pages covid service  : ", covidService)
	cache := NewCache(covidService)
	fmt.Println("creating cache : ", cache)
	p := Pages{
		Pages:        make(map[string]PageFramework),
		CovidService: covidService,
		CacheService: cache,
	}

	p.RegisterPages()
	return &p
}

//Manually register each page
func (p *Pages) RegisterPages() {
	log.Println("Registering Pages..")
	covidPage := CovidPage{}
	p.RegisterNewPage("newcovid", &covidPage)
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
