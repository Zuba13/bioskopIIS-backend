package handler

import (
	"encoding/json"
	"net/http"
	"time"

	model "bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type TheatreRepertoireHandler struct {
	TheatreRepertoireService service.TheatreRepertoireService
}

func NewTheatreRepertoireHandler(repertoireService service.TheatreRepertoireService) *TheatreRepertoireHandler {
	return &TheatreRepertoireHandler{
		TheatreRepertoireService: repertoireService,
	}
}

func (rh *TheatreRepertoireHandler) RegisterTheatreRepertoireHandler(r *mux.Router) {
	r.HandleFunc("/repertoire/timeslot", RequireRole("manager", rh.GetAvailableTimeslots)).Methods("GET")
	r.HandleFunc("/repertoire", RequireRole("manager", rh.AddProjection)).Methods("POST")
	r.HandleFunc("/repertoire", RequireRole("manager", rh.GetProjections)).Methods("GET")
	r.HandleFunc("/repertoire/consecutive-days", RequireRole("manager", rh.CountConsecutiveDaysWithProjections)).Methods("GET")
	r.HandleFunc("/repertoire/projection", RequireRole("manager", rh.CancelProjection)).Methods("DELETE")
}

func (rh *TheatreRepertoireHandler) GetAvailableTimeslots(w http.ResponseWriter, r *http.Request) {
	movieId, err := parseIntQueryParam(r, "movieId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hallId, err := parseIntQueryParam(r, "hallId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	days, err := parseIntQueryParam(r, "days", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dateStr := r.URL.Query().Get("startDate")
	// Parse the date string in the format "YYYY-MM-DD"
	startDate, err := time.Parse("2006-01-02", dateStr)
	if err != nil {
		http.Error(w, "Invalid date format", http.StatusBadRequest)
		return
	}

	if movieId == 0 || hallId == 0 || days == 0 || startDate.IsZero() {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	timeslots, err := rh.TheatreRepertoireService.GetAvailableTimeslots(movieId, hallId, days, startDate)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, timeslots)
}

func (rh *TheatreRepertoireHandler) AddProjection(w http.ResponseWriter, r *http.Request) {
	movieId, err := parseIntQueryParam(r, "movieId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hallId, err := parseIntQueryParam(r, "hallId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	days, err := parseIntQueryParam(r, "days", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ticketPrice, err := parseFloatQueryParam(r, "ticketPrice", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	timeslot := model.Timeslot{}
	err = json.NewDecoder(r.Body).Decode(&timeslot)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if movieId == 0 || hallId == 0 || days == 0 || ticketPrice == 0 {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	projections, err := rh.TheatreRepertoireService.AddProjection(movieId, hallId, days, ticketPrice, timeslot)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusCreated, projections)
}

func (rh *TheatreRepertoireHandler) GetProjections(w http.ResponseWriter, r *http.Request) {
	hallId, err := parseIntQueryParam(r, "hallId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movieId, err := parseIntQueryParam(r, "movieId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	dateStr := r.URL.Query().Get("date")
	var date time.Time

	if dateStr == "" {
		date = time.Time{}
	} else {
		// Parse the date string in the format "YYYY-MM-DD"
		date, err = time.Parse("2006-01-02", dateStr)
		if err != nil {
			http.Error(w, "Invalid date format", http.StatusBadRequest)
			return
		}
	}

	projections, err := rh.TheatreRepertoireService.GetProjections(hallId, movieId, date)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, projections)
}

func (rh *TheatreRepertoireHandler) CountConsecutiveDaysWithProjections(w http.ResponseWriter, r *http.Request) {
	projectionId, err := parseIntQueryParam(r, "projectionId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if projectionId == 0 {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	count, err := rh.TheatreRepertoireService.CountConsecutiveDaysWithProjections(uint(projectionId))
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, count)
}

func (rh *TheatreRepertoireHandler) CancelProjection(w http.ResponseWriter, r *http.Request) {
	projectionId, err := parseIntQueryParam(r, "projectionId", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	cancelConsecutive, err := parseBoolQueryParam(r, "cancelConsecutive", false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if projectionId == 0 {
		http.Error(w, "Missing required query parameters", http.StatusBadRequest)
		return
	}

	projections, err := rh.TheatreRepertoireService.CancelProjection(uint(projectionId), cancelConsecutive)
	if err != nil {
		customErr, ok := err.(*service.CustomError)
		if ok {
			http.Error(w, customErr.Message, service.ErrorCodeToHTTPStatus(customErr.Code))
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	respondWithJSON(w, http.StatusOK, projections)
}
