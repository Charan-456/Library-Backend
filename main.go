package main

import (
	"log"
	"net/http"

	"gitub.com/Charan-456/funcs/models"
	"gitub.com/Charan-456/funcs/routes"
)

var portNumber = 9002

func main() {
	models.ConnectDB()
	log.Println("Preparing server to run on port:", portNumber)
	log.Println("Loading Admin Data")
	models.DumpData()
	router := routes.Routes()
	http.ListenAndServe("0.0.0.0:9002", router)
}
