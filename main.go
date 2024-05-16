package main

import (
	"api-go-rest/database"
	"api-go-rest/models"
	"api-go-rest/routes"
	"fmt"
)

func main() {

	models.Personalidades = []models.Personalidade{
		{Id: 1, Nome: "Nome 1", Historia: "Historia 1"},
		{Id: 2, Nome: "Nome 2", Historia: "Historia 2"},
	}
	database.ConectaComBanco()
	fmt.Println("Iniciando o servidor Rest")
	routes.HandleRequest()
}

// go get -u github.com/gorilla/mux

// (docker-compose up) - subir servidor docker

// go get -u gorm.io/gorm
// go get gorm.io/driver/postgres
