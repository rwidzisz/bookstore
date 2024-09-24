package service

import "fmt"

type Book struct {
	ID         int
	Title      string
	Author     string
	Year       int
	IsBorrowed bool
}

type Library struct {
	BooksDataBase map[int]Book
}

type UserService struct {
	UsersDataBase map[int]User
}

type User struct {
	ID            int
	Name          string
	BorrowedBooks int
}

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	FindBookByTitle(title string) int
}

type UserManager interface {
	RegisterUser(user User)
	BorrowBook(library *Library, userID int, bookID int) error
	ReturnBook(library *Library, userID int, bookID int) error
}

func (library *Library) AddBook(book Book) {
	library.BooksDataBase[book.ID] = book
	fmt.Println("Book: ", book.Title, " added succesfully!")
}

func (library *Library) RemoveBook(bookID int) {
	delete(library.BooksDataBase, bookID)
	fmt.Println("Book with ID: ", bookID, " removed succesfully!")
}

func (library *Library) FindBookByTitle(title string) (int, bool) {

	for _, book := range library.BooksDataBase {
		if book.Title == title {
			fmt.Print("Book found!")
			fmt.Println(book)
			return book.ID, true
		}
	}

	fmt.Println("Book not found...")
	return 0, false
}

func (service *UserService) RegisterUser(user User) {
	service.UsersDataBase[user.ID] = user
	fmt.Println("User: ", user.Name, "registered!")
}

func (service *UserService) BorrowBook(library *Library, userID int, bookID int) error {
	user, userExists := service.UsersDataBase[userID]
	if !userExists {
		return fmt.Errorf("User with ID %d not found", userID)
	}

	book, bookExists := library.BooksDataBase[bookID]
	if !bookExists {
		return fmt.Errorf("Book with ID %d not found", bookID)
	}

	user.BorrowedBooks += 1
	service.UsersDataBase[userID] = user

	book.IsBorrowed = true
	library.BooksDataBase[bookID] = book

	fmt.Println("Book borrowed", book)
	return nil
}

func (service *UserService) ReturnBook(library *Library, userID int, bookID int) error {
	user, userExists := service.UsersDataBase[userID]
	if !userExists {
		return fmt.Errorf("User with ID %d not found", userID)
	}

	book, bookExists := library.BooksDataBase[bookID]
	if !bookExists {
		return fmt.Errorf("Book with ID %d not found", bookID)
	}

	user.BorrowedBooks -= 1
	service.UsersDataBase[userID] = user

	book.IsBorrowed = false
	library.BooksDataBase[bookID] = book

	fmt.Println("Book retuned:", book)
	return nil
}
