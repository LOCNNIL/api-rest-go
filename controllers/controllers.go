package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LOCNNIL/api-rest-go/database"
	"github.com/LOCNNIL/api-rest-go/models"
	"github.com/gorilla/mux"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Home Page")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personality;
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func ReturnOnePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r);
	id := vars["id"];
	var person models.Personality;
	database.DB.First(&person, id);
	json.NewEncoder(w).Encode(person)
}