package routes

import (
	"log"
	"net/http"

	"github.com/LOCNNIL/api-rest-go/controllers"
	"github.com/gorilla/mux"
)

func HandleRequest() {
	router := mux.NewRouter();
	router.HandleFunc("/", controllers.Home)
	router.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", router))
}
