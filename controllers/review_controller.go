package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatma0306/feedback_crud/models"
	"github.com/fatma0306/feedback_crud/utils"
	"github.com/gorilla/mux"
)

var NewReview models.Review

// CreateReview to create a review
func CreateReview(w http.ResponseWriter, r *http.Request) {
	CreateReview := &models.Review{}
	utils.ParseBody(r, CreateReview)
	b := CreateReview.CreateReview()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetReview to find all the review
func GetReview(w http.ResponseWriter, r *http.Request) {
	NewReview := models.GetAllReviews()
	res, _ := json.Marshal(NewReview)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetReviewById to find a review
func GetReviewById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewId := vars["reviewId"]
	ReviewID, err := strconv.ParseInt(reviewId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	reviewDetails, _ := models.GetReviewById(ReviewID)
	res, _ := json.Marshal(reviewDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateReview to update the review
func UpdateReview(w http.ResponseWriter, r *http.Request) {
	var updateReview = &models.Review{}
	utils.ParseBody(r, updateReview)
	vars := mux.Vars(r)
	reviewId := vars["reviewId"]
	ReviewID, err := strconv.ParseInt(reviewId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}

	reviewDetails, db := models.GetReviewById(ReviewID)

	if updateReview.ReviewContent != "" {
		reviewDetails.ReviewContent = updateReview.ReviewContent
	}

	if updateReview.ReviewState != true {
		reviewDetails.ReviewState = updateReview.ReviewState
	}

	if updateReview.CustomerID != 0 {
		reviewDetails.CustomerID = updateReview.CustomerID
	}

	db.Save(&reviewDetails)
	res, _ := json.Marshal(reviewDetails)
	w.Header().Set("Content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteReview to delete a review
func DeleteReview(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	reviewId := vars["reviewId"]
	ReviewID, err := strconv.ParseInt(reviewId, 0, 0)
	if err != nil {
		fmt.Println("Error while parsing")
	}
	review := models.DeleteReview(ReviewID)
	res, _ := json.Marshal(review)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
