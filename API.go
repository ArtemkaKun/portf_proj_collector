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
	router.HandleFunc("/stars", GetStarsCount).Methods("GET")
	router.HandleFunc("/watchers", GetProjectWatchers).Methods("GET")
	router.HandleFunc("/forks", GetProjectForks).Methods("GET")
}

func GetProjectForks(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	newMessage := new(GetProjectsRequest)
	decodeRequestMessage(req, newMessage)

	err := json.NewEncoder(writer).Encode(CalcAllForks(newMessage.Username))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetProjectWatchers(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	newMessage := new(GetProjectsRequest)
	decodeRequestMessage(req, newMessage)

	err := json.NewEncoder(writer).Encode(CalcAllWatchers(newMessage.Username))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetStarsCount(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")

	newMessage := new(GetProjectsRequest)
	decodeRequestMessage(req, newMessage)

	err := json.NewEncoder(writer).Encode(CalcAllStarts(newMessage.Username))
	if err != nil {
		EncodingJSONError(err)
	}
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
