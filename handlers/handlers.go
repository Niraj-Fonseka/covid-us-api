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

func (h *Handlers) GenerateData(w http.ResponseWriter, r *http.Request) {
	err = h.Services.Covid.GenerateNewDailyCasesData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New data generated"))
}

func (h *Handlers) GenerateCovidPage() {

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

	summaryResponse, err := h.Services.Covid.GetUSSummary()
	if err != nil {
		log.Println(err)
		return
	}

	h.Services.Graph.DrawUSMapGraph(response, summaryResponse)
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
