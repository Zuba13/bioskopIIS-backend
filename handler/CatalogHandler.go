package handler

import (
	"net/http"
	"strconv"

	"bioskop.com/projekat/bioskopIIS-backend/service"
	"github.com/gorilla/mux"
)

type MovieResponse struct {
	Title string `json:"title"`
	Year  int    `json:"year"`
	Image string `json:"image"`
}

type CatalogHandler struct {
	CatalogService  service.CatalogService
	ActorService    service.ActorService
	DirectorService service.DirectorService
}

func NewCatalogHandler(catalogService service.CatalogService, actorService service.ActorService, directorService service.DirectorService) *CatalogHandler {
	return &CatalogHandler{
		CatalogService:  catalogService,
		ActorService:    actorService,
		DirectorService: directorService,
	}
}

func (ch *CatalogHandler) RegisterCatalogHandler(r *mux.Router) {
	r.HandleFunc("/catalog/movie", RequireRole("manager", ch.GetFilteredMovies)).Methods("GET")
	r.HandleFunc("/catalog/director", RequireRole("manager", ch.GetAllDirectors)).Methods("GET")
	r.HandleFunc("/catalog/actor", RequireRole("manager", ch.GetAllActors)).Methods("GET")
	r.HandleFunc("/catalog/movie/{id:[0-9]+}", RequireRole("manager", ch.GetMovieWithAssociations)).Methods("GET")
}

func (ch *CatalogHandler) GetFilteredMovies(w http.ResponseWriter, r *http.Request) {
	title, err := parseStringQueryParam(r, "title", "")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	year, err := parseIntQueryParam(r, "year", 0)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	onlyActive, err := parseBoolQueryParam(r, "onlyActive", false)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	movies, err := ch.CatalogService.GetFilteredMovies(title, year, onlyActive)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	movieResponses := make([]MovieResponse, len(movies))
	for i, movie := range movies {
		movieResponses[i] = MovieResponse{
			Title: movie.Title,
			Year:  movie.Year,
			Image: movie.Image,
		}
	}

	respondWithJSON(w, http.StatusOK, movieResponses)
}

func (ch *CatalogHandler) GetAllDirectors(w http.ResponseWriter, r *http.Request) {
	directors, err := ch.DirectorService.GetAllDirectors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, directors)
}

func (ch *CatalogHandler) GetAllActors(w http.ResponseWriter, r *http.Request) {
	actors, err := ch.ActorService.GetAllActors()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, actors)
}

func (ch *CatalogHandler) GetMovieWithAssociations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid movie ID", http.StatusBadRequest)
		return
	}

	movie, err := ch.CatalogService.GetMovieWithAssociations(uint(id))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, movie)
}

func parseStringQueryParam(r *http.Request, name string, defaultValue string) (string, error) {
	value := r.URL.Query().Get(name)
	if value == "" {
		return defaultValue, nil
	}
	return value, nil
}

func parseIntQueryParam(r *http.Request, name string, defaultValue int) (int, error) {
	valueStr := r.URL.Query().Get(name)
	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.Atoi(valueStr)
}

func parseBoolQueryParam(r *http.Request, name string, defaultValue bool) (bool, error) {
	valueStr := r.URL.Query().Get(name)
	if valueStr == "" {
		return defaultValue, nil
	}
	return strconv.ParseBool(valueStr)
}
