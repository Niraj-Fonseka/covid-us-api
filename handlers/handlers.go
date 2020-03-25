package handlers

import (
	"covid-us-api/services"
	"log"
	"net/http"
)

type Handlers struct {
	Services *services.Services
	Pages    *services.Pages
}

func RegisterHandlers(s *services.Services, p *services.Pages) *Handlers {
	return &Handlers{
		Services: s,
		Pages:    p,
	}
}

func (h *Handlers) GenerateDailyData(w http.ResponseWriter, r *http.Request) {
	err := h.Services.Covid.GenerateNewDailyCasesData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New data generated"))
}

func (h *Handlers) GenerateSummaryData(w http.ResponseWriter, r *http.Request) {
	err := h.Services.Covid.GenerateNewOverallCasesData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New data generated"))
}

func (h *Handlers) GenerateCountyData(w http.ResponseWriter, r *http.Request) {
	err := h.Services.Covid.GenerateUSCountyData()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte("New data generated"))
}

func (h *Handlers) UploadMainPage(w http.ResponseWriter, r *http.Request) {
	log.Println("Uploading main page")
	h.Services.Covid.UploadMainPage()
}
func (h *Handlers) UploadStatePages(w http.ResponseWriter, r *http.Request) {
	h.Services.Covid.UploadAllStateFiles()
}

func (h *Handlers) RenderPage(w http.ResponseWriter, r *http.Request) {

	pageParam := r.URL.Query().Get("page")
	if pageParam == "" {
		log.Println("rendering everything")
	} else {
		log.Printf("Rendering just this page : %s\n", pageParam)
	}
	h.Pages.RenderPages(pageParam)

}
