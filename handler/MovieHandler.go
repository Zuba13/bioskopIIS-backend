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

type MovieHandler struct {
	MovieService service.MovieService
}

func NewMovieHandler(movieService service.MovieService) *MovieHandler {
	return &MovieHandler{
		MovieService: movieService,
	}
}

func (mh *MovieHandler) RegisterMovieHandler(r *mux.Router) {
	r.HandleFunc("/movies", mh.CreateMovie).Methods("POST")
	r.HandleFunc("/movies", mh.GetAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", mh.GetMovieByID).Methods("GET")
	r.HandleFunc("/movies/{id}", mh.UpdateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", mh.DeleteMovie).Methods("DELETE")
}

func (mh *MovieHandler) CreateMovie(w http.ResponseWriter, r *http.Request) {
	var newMovie model.Movie
	err := json.NewDecoder(r.Body).Decode(&newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set current time for CreatedAt and UpdatedAt fields
	newMovie.CreatedAt = time.Now()
	newMovie.UpdatedAt = time.Now()

	movie, err := mh.MovieService.CreateMovie(newMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, movie)
}

func (mh *MovieHandler) GetAllMovies(w http.ResponseWriter, r *http.Request) {
	movies, err := mh.MovieService.GetAllMovies()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, movies)
}

func (mh *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movie, err := mh.MovieService.GetMovieByID(uint(movieID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

func (mh *MovieHandler) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedMovie model.Movie
	err = json.NewDecoder(r.Body).Decode(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedMovie.ID = uint(movieID)
	updatedMovie.UpdatedAt = time.Now()

	err = mh.MovieService.UpdateMovie(&updatedMovie)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedMovie)
}

func (mh *MovieHandler) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = mh.MovieService.DeleteMovie(uint(movieID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Movie deleted successfully"})
}

