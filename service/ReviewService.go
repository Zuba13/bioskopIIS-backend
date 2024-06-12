package service

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"bioskop.com/projekat/bioskopIIS-backend/repo"
)

// ReviewService handles operations related to reviews.
type ReviewService struct {
	ReviewRepo repo.ReviewRepository
}

// NewReviewService creates a new instance of the ReviewService.
func NewReviewService(reviewRepo repo.ReviewRepository) *ReviewService {
	return &ReviewService{
		ReviewRepo: reviewRepo,
	}
}

// CreateReview creates a new review.
func (rs *ReviewService) CreateReview(review model.Review) (*model.Review, error) {
	return rs.ReviewRepo.CreateReview(review)
}

// GetAllReviews retrieves all reviews.
func (rs *ReviewService) GetAllReviews() ([]model.Review, error) {
	return rs.ReviewRepo.GetAllReviews()
}

// GetReviewByID retrieves a review by its ID.
func (rs *ReviewService) GetReviewByID(id uint) (*model.Review, error) {
	return rs.ReviewRepo.GetReviewByID(id)
}

// UpdateReview updates an existing review.
func (rs *ReviewService) UpdateReview(review *model.Review) error {
	return rs.ReviewRepo.UpdateReview(review)
}

// DeleteReview deletes a review by its ID.
func (rs *ReviewService) DeleteReview(id uint) error {
	return rs.ReviewRepo.DeleteReview(id)
}

// GetReviewsByUserID retrieves reviews by user ID.
func (rs *ReviewService) GetReviewsByUserID(userID uint) ([]model.Review, error) {
	return rs.ReviewRepo.GetReviewsByUserID(userID)
}

// GetReviewsByMovieID retrieves reviews by movie ID.
func (rs *ReviewService) GetReviewsByMovieID(movieID uint) ([]model.Review, error) {
	return rs.ReviewRepo.GetReviewsByMovieID(movieID)
}
