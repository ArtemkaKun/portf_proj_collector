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
		projects := GetLastActiveProjects(uint16(ParameterToInt(params["count"], writer)), params["username"])
		err = json.NewEncoder(writer).Encode(projects)
	}

	if err != nil {
		EncodingJSONError(err, writer)
	}
}

func ParameterToInt(param string, writer http.ResponseWriter) (projCount int) {
	projCount, err := strconv.Atoi(param)

	if err != nil {
		errorRequest := map[string]string{}
		errorMessage := fmt.Sprintf("%v", fmt.Errorf("Cannot convert 'count' parameter to number: %v", err))
		errorRequest["error"] = errorMessage

		err = json.NewEncoder(writer).Encode(errorRequest)
		if err != nil {
			EncodingJSONError(err, writer)
		}

		fmt.Println(errorMessage)
	}
	return
}

func EncodingJSONError(err error, writer http.ResponseWriter) {
	errorRequest := map[string]string{}
	errorMessage := fmt.Sprintf("%v", fmt.Errorf("Error while decoding JSON: %v\n", err))
	errorRequest["error"] = errorMessage

	json.NewEncoder(writer).Encode(errorRequest)
	fmt.Println(errorMessage)
}
