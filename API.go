package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/activeProjects", GetActiveProjects).Methods("GET")
}

func GetActiveProjects(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	var newMessage GetProjectsRequest
	decodeRequestMessage(req, &newMessage)

	projects := GetLastActiveProjects(newMessage.ProjectsCount, newMessage.Username)

	err := json.NewEncoder(writer).Encode(projects)
	if err != nil {
		EncodingJSONError(err)
	}
}

func decodeRequestMessage(req *http.Request, newMessage *GetProjectsRequest) {
	err := json.NewDecoder(req.Body).Decode(newMessage)
	DecodingJSONError(err)
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}

func DecodingJSONError(err error) {
	if err != nil {
		fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
	}
}
