package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/fatma0306/feedback_crud/config"
	"github.com/fatma0306/feedback_crud/models"
	"github.com/fatma0306/feedback_crud/routes"
	"github.com/gorilla/mux"
)

func Run() {
	db := config.Connect()
	if !db.HasTable(&models.Review{}) {
		db.Debug().CreateTable(&models.Review{})
	}
	db.Close()
	listen(8000)
}

func listen(l int) {
	port := fmt.Sprintf(":%d", l)
	fmt.Printf("Listening Port %s...\n", port)
	r := mux.NewRouter()
	routes.RegisterReviewRoutes(r)
	log.Fatal(http.ListenAndServe(port, r))
}

func main() {
	Run()
}

/*func main() {

	r := mux.NewRouter()
	routes.RegisterReviewRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("localhost:3000", r))
}*/
