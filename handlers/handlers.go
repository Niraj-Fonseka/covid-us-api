package handlers

import (
	"covid-us-api/services"
	"log"
	"net/http"
)

type Handlers struct {
	Services *services.Services
}

func RegisterHandlers(s *services.Services) *Handlers {
	return &Handlers{
		Services: s,
	}
}

func (h *Handlers) SlackHandler(w http.ResponseWriter, r *http.Request) {
	h.Services.Covid.GetDailyCasesUS()
}

func (h *Handlers) DrawGraph(w http.ResponseWriter, r *http.Request) {
	response, err := h.Services.Covid.GetDailyCasesUS()
	if err != nil {
		log.Println(err)
		return
	}
	h.Services.Graph.DrawDeathsGraph(response)
}

func (h *Handlers) DrawGraphUSMAP(w http.ResponseWriter, r *http.Request) {
	response, err := h.Services.Covid.GetDailyCasesUS()
	if err != nil {
		log.Println(err)
		return
	}
	h.Services.Graph.DrawUSMapGraph(response)
}

func (h *Handlers) DrawGraphState(w http.ResponseWriter, r *http.Request) {
	queryValues := r.URL.Query()
	state := queryValues.Get("state")
	response, err := h.Services.Covid.GetDailyCasesUSByState(state)
	if err != nil {
		log.Println(err)
		return
	}
	h.Services.Graph.RenderStatePage(state, response)
}
