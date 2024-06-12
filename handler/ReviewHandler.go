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

type ReviewHandler struct {
	ReviewService service.ReviewService
}

func NewReviewHandler(reviewService service.ReviewService) *ReviewHandler {
	return &ReviewHandler{
		ReviewService: reviewService,
	}
}

func (rh *ReviewHandler) RegisterReviewHandler(r *mux.Router) {
	r.HandleFunc("/reviews", rh.CreateReview).Methods("POST")
	r.HandleFunc("/reviews", rh.GetAllReviews).Methods("GET")
	r.HandleFunc("/reviews/{id}", rh.GetReviewByID).Methods("GET")
	r.HandleFunc("/reviews/{id}", rh.UpdateReview).Methods("PUT")
	r.HandleFunc("/reviews/{id}", rh.DeleteReview).Methods("DELETE")
	r.HandleFunc("/reviews/user/{id}", rh.GetReviewsByUserID).Methods("GET")
	r.HandleFunc("/reviews/movie/{id}", rh.GetReviewsByMovieID).Methods("GET")
}

func (rh *ReviewHandler) CreateReview(w http.ResponseWriter, r *http.Request) {
	var newReview model.Review
	err := json.NewDecoder(r.Body).Decode(&newReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set current time for CreatedAt and UpdatedAt fields
	newReview.CreatedAt = time.Now()
	newReview.UpdatedAt = time.Now()

	review, err := rh.ReviewService.CreateReview(newReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusCreated, review)
}

func (rh *ReviewHandler) GetAllReviews(w http.ResponseWriter, r *http.Request) {
	reviews, err := rh.ReviewService.GetAllReviews()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, reviews)
}

func (rh *ReviewHandler) GetReviewByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	review, err := rh.ReviewService.GetReviewByID(uint(reviewID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	respondWithJSON(w, http.StatusOK, review)
}

func (rh *ReviewHandler) UpdateReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var updatedReview model.Review
	err = json.NewDecoder(r.Body).Decode(&updatedReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	updatedReview.ID = uint(reviewID)
	updatedReview.UpdatedAt = time.Now()

	err = rh.ReviewService.UpdateReview(&updatedReview)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, updatedReview)
}

func (rh *ReviewHandler) DeleteReview(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	reviewID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = rh.ReviewService.DeleteReview(uint(reviewID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, map[string]string{"message": "Review deleted successfully"})
}

func (rh *ReviewHandler) GetReviewsByUserID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reviews, err := rh.ReviewService.GetReviewsByUserID(uint(userID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, reviews)
}

func (rh *ReviewHandler) GetReviewsByMovieID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	movieID, err := strconv.ParseUint(params["id"], 10, 64)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	reviews, err := rh.ReviewService.GetReviewsByMovieID(uint(movieID))
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	respondWithJSON(w, http.StatusOK, reviews)
}
