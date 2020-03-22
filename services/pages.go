package services

import "log"

type Pages struct {
	Pages map[string]PageFramework
}

func RegisterPages() *Pages {
	return &Pages{
		Pages: make(map[string]PageFramework),
	}
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
