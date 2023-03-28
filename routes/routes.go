package routes

import (
	"log"
	"net/http"

	"github.com/LOCNNIL/api-rest/go/controllers"
)

func HandleRequest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalities", controllers.AllPersonalities)
	log.Fatal(http.ListenAndServe(":8000", nil))

}
