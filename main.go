package main

import (
	"fmt"

	"github.com/LOCNNIL/api-rest-go/models"
	"github.com/LOCNNIL/api-rest-go/routes"
)

func main() {
	models.Personalities = []models.Personality{
		{Name: "First Name", History: "First History"},
		{Name: "Secound Name", History: "Secound History"},
	}
	fmt.Println("Starting the REST server with GO!")
	routes.HandleRequest()
}
