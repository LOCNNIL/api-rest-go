package main

import (
	"fmt"

	"github.com/LOCNNIL/api-rest-go/models"
	"github.com/LOCNNIL/api-rest-go/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "1º Name", History: "First History"},
		{Name: "2º Name", History: "Secound History"},
		{Name: "3º Name", History: "Third History"},
	}
	fmt.Println("Starting the REST server with GO!")
	routes.HandleRequest()
}
