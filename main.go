package main

import (
	"fmt"

	"github.com/LOCNNIL/api-rest-go/database"
	"github.com/LOCNNIL/api-rest-go/models"
	"github.com/LOCNNIL/api-rest-go/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Id: 1, Name: "1º Name", History: "First History"},
		{Id: 2, Name: "2º Name", History: "Secound History"},
		{Id: 3, Name: "3º Name", History: "Third History"},
	}

	database.ConectWithDatabase()
	fmt.Println("Starting the REST server with GO!")
	routes.HandleRequest()
}
