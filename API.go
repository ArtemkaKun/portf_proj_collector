package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var router = mux.NewRouter()

func init() {
	router.HandleFunc("/activeProjects/{count}/{username}", GetActiveProjects).Methods("GET")
	router.HandleFunc("/stars/{username}", GetStarsCount).Methods("GET")
	router.HandleFunc("/watchers/{username}", GetProjectWatchers).Methods("GET")
	router.HandleFunc("/forks/{username}", GetProjectForks).Methods("GET")
}

func GetProjectForks(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	err := json.NewEncoder(writer).Encode(CalcAllForks(params["username"]))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetProjectWatchers(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	err := json.NewEncoder(writer).Encode(CalcAllWatchers(params["username"]))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetStarsCount(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	err := json.NewEncoder(writer).Encode(CalcAllStarts(params["username"]))
	if err != nil {
		EncodingJSONError(err)
	}
}

func GetActiveProjects(writer http.ResponseWriter, req *http.Request) {
	writer.Header().Set("Content-Type", "application/json")
	params := mux.Vars(req)

	projCount, err := strconv.Atoi(params["count"])
	if err != nil {
		fmt.Println(fmt.Errorf("Cannot convert 'count' parameter to number: %v", err))
	}

	projects := GetLastActiveProjects(uint16(projCount), params["username"])

	err = json.NewEncoder(writer).Encode(projects)
	if err != nil {
		EncodingJSONError(err)
	}
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}
