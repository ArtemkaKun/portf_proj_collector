package main

import (
	"github.com/rs/cors"
	"log"
	"net/http"
)

func main() {
	log.Fatal(http.ListenAndServe(":8001", cors.Default().Handler(Router)))
}
