package main

import (
	"forum/database"
	"forum/errorhandle"
	"forum/router"
	"forum/templ"
	"log"
	"net/http"
	"os"
)

func init() {
	database.SetUp()
	templ.SetUp()
}
func main() {
	defer database.Close()

	mux := router.ServeMux()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("Server is listening:", port)
	errorhandle.CheckErr(http.ListenAndServe(":"+port, mux))
}
