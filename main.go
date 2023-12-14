package main

import (
	"library/routes"
	"library/services"
	"library/utility"
	"log"
	"net/http"
)

func main() {

	var db = utility.GetConnection()
	services.SetDB(db)
	var appRouter = routes.CreateRouter()

	log.Println("Listening on Port 8080")
	log.Fatal(http.ListenAndServe(":8080", appRouter))

}
