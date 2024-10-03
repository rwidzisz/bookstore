package main

import (
	"bookstore/router"
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Rest API with connection to MongoDB:")
	r := router.Router()
	fmt.Println("Server is getting started...")
	log.Fatal(http.ListenAndServe(":10000", r))
	fmt.Println("Listening at port: 10000")

}
