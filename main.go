package main

import (
	"log"
	"net/http"

	"github.com/Charan-456/Library-Backend/models"
	"github.com/Charan-456/Library-Backend/routes"
	// "github.com/joho/godotenv"
)

var portNumber = 9002

func main() {
	// errors := godotenv.Load()
	// if errors != nil {
	// 	log.Fatal("Error loading .env file")
	// }
	models.ConnectDB()
	log.Println("Preparing server to run on port:", portNumber)
	log.Println("Loading Admin Data")
	models.DumpData()
	router := routes.Routes()
	http.ListenAndServe("0.0.0.0:9002", router)
}
