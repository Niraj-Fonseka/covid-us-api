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
	h.Services.Graph.DrawGraphTwo(response)
}
