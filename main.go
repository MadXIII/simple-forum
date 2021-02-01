package main

import (
	"forum/database"
	"forum/errorhandle"
	"forum/front"
	"log"
	"net/http"
	"os"
)

func init() {
	database.SetUp()
	front.SetUp()
}
func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8888"
	}

	log.Println("Server is listening:", port)
	errorhandle.CheckErr(http.ListenAndServe(":"+port, nil))
}
