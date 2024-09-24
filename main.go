package main

import (
	"fmt"
	"methods/service"
)

func main() {
	bk1 := service.Book{ID: 1, Title: "Witcher", Author: "Sapkowski", Year: 1992, IsBorrowed: false}
	bk2 := service.Book{ID: 2, Title: "Lord Of The Rings", Author: "Tolkien", Year: 1960, IsBorrowed: false}
	bk3 := service.Book{ID: 3, Title: "Mythology", Author: "Parandowski", Year: 1920, IsBorrowed: false}

	//Dodawanie ksiazek
	MainLibrary := service.Library{BooksDataBase: make(map[int]service.Book)}
	MainLibrary.AddBook(bk1)
	MainLibrary.AddBook(bk2)
	MainLibrary.AddBook(bk3)

	usr1 := service.User{ID: 1, Name: "Rafał", BorrowedBooks: 0}
	usr2 := service.User{ID: 2, Name: "Felix", BorrowedBooks: 0}

	//Dodawanie userów
	MainUserService := service.UserService{UsersDataBase: make(map[int]service.User)}
	MainUserService.RegisterUser(usr1)
	MainUserService.RegisterUser(usr2)

	fmt.Println(MainLibrary)
	//Szukanie ksiazek, pozyczanie i oddawanie
	bookID, bookExists := MainLibrary.FindBookByTitle("Witcher")
	if bookExists {
		MainUserService.BorrowBook(&MainLibrary, 1, bookID)
		MainUserService.ReturnBook(&MainLibrary, 1, bookID)
	}

	fmt.Println(MainLibrary)
}
