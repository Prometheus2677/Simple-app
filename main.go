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

	log.Println("Listening on Port 8000")
	log.Fatal(http.ListenAndServe(":8000", appRouter))

}
