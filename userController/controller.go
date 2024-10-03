package userController

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

type UserManager interface {
	RegisterUser(user User)
	BorrowBook(userID int, bookID int) error
	ReturnBook(userID int, bookID int) error
}

const connectionString = "mongodb+srv://rwidzisz:papuga33@restapitesting.5tdayhy.mongodb.net/?retryWrites=true&w=majority&appName=RestApiTesting"
const databaseName = "bookstore"

var collection *mongo.Collection

func init() {
	collection = getTheCollection("users")
	fmt.Println("Collection:", collection.Name(), " is ready!")
}

func getTheCollection(collectionName string) *mongo.Collection {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}

	return client.Database(databaseName).Collection(collectionName)
}

func registerUser(user User) {
	insertion, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("User: ", user.Name, "registered succesfully with ID: ", insertion.InsertedID)
}

func rentingBook(userID string, bookID string, option string) {
	borrow := true
	incrementation := 1

	if option == "return" {
		borrow = false
		incrementation = -1
	}

	collection = getTheCollection("books")

	id, _ := primitive.ObjectIDFromHex(bookID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isborrowed": borrow}}

	_, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	collection = getTheCollection("users")

	id, _ = primitive.ObjectIDFromHex(userID)
	filter = bson.M{"_id": id}
	update = bson.M{"$inc": bson.M{"borrowedbooks": incrementation}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Book:", bookID, " has been ", option, "ed by user:", userID)
}

func RegisterUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "POST")

	var user User
	_ = json.NewDecoder(r.Body).Decode(&user)
	registerUser(user)

	json.NewEncoder(w).Encode(user)

}

func RentBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "appplication/x-www-form-urlencode")
	w.Header().Set("Allow-Control-Allow-Methods", "PUT")

	parameters := mux.Vars(r)
	userID := parameters["userId"]
	bookID := parameters["bookId"]
	option := parameters["option"]
	rentingBook(userID, bookID, option)

}
