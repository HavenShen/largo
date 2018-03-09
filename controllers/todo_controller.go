package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/HavenShen/largo/models"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Welcome!\n")
}

func TodoIndex(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(models.GetTodos()); err != nil {
		panic(err)
	}
}
