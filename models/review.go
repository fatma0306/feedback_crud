package models

import (
	"github.com/fatma0306/feedback_crud/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Review struct {
	gorm.Model
	//ReviewID      int    `json:"ReviewID"`
	ReviewContent string `json:"ReviewContent"`
	ReviewState   bool   `json:"ReviewState"`
	CustomerID    int    `json:"CustomerID"`
}

func init() {
	db = config.Connect()
	//db = config.GetDB()
	db.AutoMigrate(&Review{})
}

// CreateReview to create a review
func (b *Review) CreateReview() *Review {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// GetAllReviews to get all the reviews
func GetAllReviews() []Review {
	var Reviews []Review
	db.Find(&Reviews)
	return Reviews
}

// GetReviewById to get a reviews
func GetReviewById(ReviewID int64) (*Review, *gorm.DB) {
	var getReview Review
	db := db.Where("ReviewID = ?", ReviewID).Find(&getReview)
	return &getReview, db

}

// DeleteReview to delete a review
func DeleteReview(ReviewID int64) Review {
	var review Review
	db.Where("ReviewID = ?", ReviewID).Delete(review)
	return review

}
