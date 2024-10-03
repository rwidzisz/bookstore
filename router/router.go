package router

import (
	"bookstore/libraryController"
	"bookstore/userController"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func Router() *mux.Router {
	myRouter := mux.NewRouter()

	myRouter.HandleFunc("/", homePage)

	myRouter.HandleFunc("/books", libraryController.GetAllBooks).Methods("GET")
	myRouter.HandleFunc("/book", libraryController.InsertOneBook).Methods("POST")
	myRouter.HandleFunc("/book/{id}", libraryController.RemoveOneBook).Methods("DELETE")
	myRouter.HandleFunc("/book/{id}", libraryController.MarkAsBorrowed).Methods("PUT")
	myRouter.HandleFunc("/book/{title}", libraryController.FindBookByTitle).Methods("GET")

	myRouter.HandleFunc("/user", userController.RegisterUser).Methods("POST")
	myRouter.HandleFunc("/user/{userId}/{option}/{bookId}", userController.RentBook).Methods("PUT")

	//dorobić komiksy

	log.Fatal(http.ListenAndServe(":10000", myRouter))

	return myRouter
}

func homePage(writer http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(writer, "Welcome to the Homepage!")
}
