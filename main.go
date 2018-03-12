package main

import (
	"github.com/HavenShen/largo/config"
	"github.com/HavenShen/largo/routes"
	"log"
	"net/http"
)

func main() {
	router := routes.NewRouter()
	log.Fatal(http.ListenAndServe(":8888", router))
}
