package handler

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type ProjectionHandler struct {
	ProjectionService service.ProjectionService
}

func NewProjectionHandler(projectionService service.ProjectionService) *ProjectionHandler {
	return &ProjectionHandler{
		ProjectionService: projectionService,
	}
}

func (ph *ProjectionHandler) RegisterProjectionHandler(r *mux.Router) {
	r.HandleFunc("/projections", ph.CreateProjection).Methods("POST")
	r.HandleFunc("/projections", ph.GetAllProjections).Methods("GET")
	r.HandleFunc("/projections/{id}", ph.GetProjectionByID).Methods("GET")
	r.HandleFunc("/projections/{id}", ph.UpdateProjection).Methods("PUT")
	r.HandleFunc("/projections/{id}", ph.DeleteProjection).Methods("DELETE")
}

func (ph *ProjectionHandler) CreateProjection(w http.ResponseWriter, r *http.Request) {
	var newProjection model.Projection
	err := json.NewDecoder(r.Body).Decode(&newProjection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set current time for CreatedAt and UpdatedAt fields
	newProjection.CreatedAt = time.Now()
	newProjection.UpdatedAt = time.Now()

	projection, err := ph.ProjectionService.CreateProjection(newProjection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, projection)
}

func (ph *ProjectionHandler) GetAllProjections(w http.ResponseWriter, r *http.Request) {
	projections, err := ph.ProjectionService.GetAllProjections()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, projections)
}

func (ph *ProjectionHandler) GetProjectionByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectionID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	projection, err := ph.ProjectionService.GetProjectionByID(uint(projectionID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, projection)
}

func (ph *ProjectionHandler) UpdateProjection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectionID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedProjection model.Projection
	err = json.NewDecoder(r.Body).Decode(&updatedProjection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedProjection.ID = uint(projectionID)
	updatedProjection.UpdatedAt = time.Now()

	err = ph.ProjectionService.UpdateProjection(&updatedProjection)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedProjection)
}

func (ph *ProjectionHandler) DeleteProjection(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	projectionID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ph.ProjectionService.DeleteProjection(uint(projectionID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Projection deleted successfully"})
}
