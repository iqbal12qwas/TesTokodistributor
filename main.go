package main

import (
	"log"
	"main/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// fmt.Println("Hello, World!")

	log.Println("Starting the HTTP server on port 8090")

	router := mux.NewRouter().StrictSlash(true)
	initaliseHandlers(router)
	log.Fatal(http.ListenAndServe(":8090", router))
}

func initaliseHandlers(router *mux.Router) {
	// Route GET By Params ID "Task 1"
	router.HandleFunc("/user/{id}", controllers.GetUrlParamUserByID).Methods("GET")
	router.HandleFunc("/user", controllers.GetReqParamUserByID).Methods("GET")

	// Route GET By Params ID "Task 2"
	router.HandleFunc("/bowling", controllers.ChooseContainer).Methods("GET")
}
