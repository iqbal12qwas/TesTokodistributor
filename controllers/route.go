package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

//Model / Struct User
type User struct {
	Id        int    `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Age       int    `json:"age"`
}

//Model / Struct Response
type Response struct {
	Http    string `json:"http"`
	Message string `json:"message"`
	Data    []User `json:"data"`
}

func GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Get Params By ID User
	vars := mux.Vars(r)
	key := vars["id"]

	// Data statis Get
	var user []User

	if key == "1" { //If key/id = 1
		user = append(user, User{Id: 1, FirstName: "Muhammad", LastName: "Iqbal Zulkifli", Age: 23})
	} else { //If key/id != 1
		user = nil
	}

	// Show Response & Parse To Json
	w.Header().Set("Content-Type", "application/json")
	resp := Response{Http: "GET/200", Message: "Success", Data: user}
	json.NewEncoder(w).Encode(resp)
}
