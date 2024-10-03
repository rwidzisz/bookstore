package libraryController

import "go.mongodb.org/mongo-driver/bson/primitive"

type Book struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	Title      string             `bson:"title,omitempty"`
	Author     string             `bson:"author,omitempty"`
	Year       int32              `bson:"year,omitempty"`
	IsBorrowed bool               `bson:"isborrowed"`
}

//Note: The omitempty means that if there is no data in the particular field,
//when saved to MongoDB the field will not exist on the document rather than
//existing with an empty value.
