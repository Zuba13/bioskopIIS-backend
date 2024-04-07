package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type HallHandler struct {
	HallService service.HallService
}

func NewHallHandler(hallService service.HallService) *HallHandler {
	return &HallHandler{
		HallService: hallService,
	}
}

func (hh *HallHandler) RegisterHallHandler(r *mux.Router) {
	r.HandleFunc("/halls", hh.CreateHall).Methods("POST")
	r.HandleFunc("/halls", hh.GetAllHalls).Methods("GET")
	r.HandleFunc("/halls/{id}", hh.GetHallByID).Methods("GET")
	r.HandleFunc("/halls/{id}", hh.UpdateHall).Methods("PUT")
	r.HandleFunc("/halls/{id}", hh.DeleteHall).Methods("DELETE")
}

func (hh *HallHandler) CreateHall(w http.ResponseWriter, r *http.Request) {
	var newHall model.Hall
	err := json.NewDecoder(r.Body).Decode(&newHall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hall, err := hh.HallService.CreateHall(newHall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, hall)
}

func (hh *HallHandler) GetAllHalls(w http.ResponseWriter, r *http.Request) {
	halls, err := hh.HallService.GetAllHalls()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, halls)
}

func (hh *HallHandler) GetHallByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hallID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	hall, err := hh.HallService.GetHallByID(uint(hallID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, hall)
}

func (hh *HallHandler) UpdateHall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hallID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedHall model.Hall
	err = json.NewDecoder(r.Body).Decode(&updatedHall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedHall.ID = uint(hallID)
	err = hh.HallService.UpdateHall(&updatedHall)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedHall)
}

func (hh *HallHandler) DeleteHall(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	hallID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = hh.HallService.DeleteHall(uint(hallID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Hall deleted successfully"})
}
