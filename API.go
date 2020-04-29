package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var Router = mux.NewRouter()

func init() {
	Router.HandleFunc("/activeProjects/{count}/{username}", GetActiveProjects).Methods("GET")
	Router.HandleFunc("/stars/{username}", GetStarsCount).Methods("GET")
	Router.HandleFunc("/watchers/{username}", GetProjectWatchers).Methods("GET")
	Router.HandleFunc("/forks/{username}", GetProjectForks).Methods("GET")
}

func GetProjectForks(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	WriteSimpleStatsAnswer(0, writer, params)
}

func GetProjectWatchers(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	WriteSimpleStatsAnswer(1, writer, params)
}

func GetStarsCount(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)

	WriteSimpleStatsAnswer(2, writer, params)
}

func GetActiveProjects(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	WriteSimpleStatsAnswer(3, writer, params)
}

func WriteSimpleStatsAnswer(answerType uint8, writer http.ResponseWriter, params map[string]string) {
	writer.Header().Set("Content-Type", "application/json")
	var err error

	switch answerType {
	case 0:
		err = json.NewEncoder(writer).Encode(CalcAllForks(params["username"]))
	case 1:
		err = json.NewEncoder(writer).Encode(CalcAllWatchers(params["username"]))
	case 2:
		err = json.NewEncoder(writer).Encode(CalcAllStarts(params["username"]))
	case 3:
		projects := GetLastActiveProjects(uint16(ParameterToInt(params)), params["username"])
		err = json.NewEncoder(writer).Encode(projects)
	}

	if err != nil {
		EncodingJSONError(err)
	}
}

func ParameterToInt(params map[string]string) (projCount int) {
	projCount, err := strconv.Atoi(params["count"])
	if err != nil {
		fmt.Println(fmt.Errorf("Cannot convert 'count' parameter to number: %v", err))
	}
	return
}

func EncodingJSONError(err error) {
	fmt.Println(fmt.Errorf("Error while decoding JSON: %v\n", err))
}
