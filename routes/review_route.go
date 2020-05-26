package routes

import (
	"github.com/fatma0306/feedback_crud/controllers"
	"github.com/gorilla/mux"
)

// RegisterReviewRoutes
var RegisterReviewRoutes = func(router *mux.Router) {

	router.HandleFunc("/review", controllers.CreateReview).Methods("POST")
	router.HandleFunc("/review", controllers.GetReview).Methods("GET")
	router.HandleFunc("/review{reviewId}", controllers.GetReviewById).Methods("GET")
	router.HandleFunc("/review{reviewId}", controllers.UpdateReview).Methods("PUT")
	router.HandleFunc("/review{reviewId}", controllers.DeleteReview).Methods("DELETE")

}
