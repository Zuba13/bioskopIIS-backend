package repo

import (
	"bioskop.com/projekat/bioskopIIS-backend/model"
	"gorm.io/gorm"
)

type ReviewRepository struct {
	DatabaseConnection *gorm.DB
}

func NewReviewRepository(db *gorm.DB) *ReviewRepository {
	return &ReviewRepository{DatabaseConnection: db}
}

func (rr *ReviewRepository) CreateReview(review model.Review) (*model.Review, error) {
	if err := rr.DatabaseConnection.Create(&review).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (rr *ReviewRepository) GetAllReviews() ([]model.Review, error) {
	var reviews []model.Review
	if err := rr.DatabaseConnection.Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (rr *ReviewRepository) GetReviewByID(id uint) (*model.Review, error) {
	var review model.Review
	if err := rr.DatabaseConnection.First(&review, id).Error; err != nil {
		return nil, err
	}
	return &review, nil
}

func (rr *ReviewRepository) UpdateReview(review *model.Review) error {
	return rr.DatabaseConnection.Save(review).Error
}

func (rr *ReviewRepository) DeleteReview(id uint) error {
	return rr.DatabaseConnection.Delete(&model.Review{}, id).Error
}

func (rr *ReviewRepository) GetReviewsByUserID(userID uint) ([]model.Review, error) {
	var reviews []model.Review
	if err := rr.DatabaseConnection.Where("user_id = ?", userID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}

func (rr *ReviewRepository) GetReviewsByMovieID(movieID uint) ([]model.Review, error) {
	var reviews []model.Review
	if err := rr.DatabaseConnection.Where("movie_id = ?", movieID).Find(&reviews).Error; err != nil {
		return nil, err
	}
	return reviews, nil
}
