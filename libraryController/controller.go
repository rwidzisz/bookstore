package libraryController

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type LibraryManager interface {
	AddBook(book Book)
	RemoveBook(bookID int)
	FindBookByTitle(title string) int
}

const connectionString = "mongodb+srv://rwidzisz:papuga33@restapitesting.5tdayhy.mongodb.net/?retryWrites=true&w=majority&appName=RestApiTesting"
const databaseName = "bookstore"
const collectionName = "books"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	collection = client.Database(databaseName).Collection(collectionName)
	fmt.Println("Collection:", collectionName, " is ready!")
}

func insertOneBook(book Book) {
	insertion, err := collection.InsertOne(context.Background(), book)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Book named:", book.Title, "inserted succesfuly with ID: ", insertion.InsertedID)
}

func removeOneBook(bookID string) {
	// ObjectIDFromHex konwertuje stringa w primityw Mongo
	id, _ := primitive.ObjectIDFromHex(string(bookID))
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book with ID: ", bookID, " removed succesfully with count: ", deleteCount)
}

func markAsBorrowed(bookID string) {
	id, _ := primitive.ObjectIDFromHex(string(bookID))
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isborrowed": true}}

	updateCount, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book with ID: ", bookID, " updated succesfully with count: ", updateCount)

}

func getAllBooks() []Book {
	// Tworzenie kursora danych z kolekcji po którym moemy się iterować
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var books []Book

	for cursor.Next(context.Background()) {
		var book Book
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, book)
	}

	defer cursor.Close(context.Background())
	return books
}

func findBookByTitle(title string) Book {
	cursor, err := collection.Find(context.Background(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	var searchedBook Book

	for cursor.Next(context.Background()) {
		var book Book
		err := cursor.Decode(&book)
		if err != nil {
			log.Fatal(err)
		}
		if book.Title == title {
			searchedBook = book
			break
		}
	}

	defer cursor.Close(context.Background())
	return searchedBook
}

func InsertOneBook(w http.ResponseWriter, r *http.Request) {
	//Headesry definiują między innymi jakie mogą działać

	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var book Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatal(err)
	}
	insertOneBook(book)

	json.NewEncoder(w).Encode(book)

}

func RemoveOneBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "DELETE")

	parameters := mux.Vars(r)
	removeOneBook(parameters["id"])

	json.NewEncoder(w).Encode(parameters["id"])

}

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")

	allBooks := getAllBooks()

	json.NewEncoder(w).Encode(allBooks)
}

func FindBookByTitle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")

	parameters := mux.Vars(r)
	book := findBookByTitle(parameters["title"])

	json.NewEncoder(w).Encode(book)
}

func MarkAsBorrowed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	parameters := mux.Vars(r)
	markAsBorrowed(parameters["id"])

	json.NewEncoder(w).Encode(parameters["id"])

}
