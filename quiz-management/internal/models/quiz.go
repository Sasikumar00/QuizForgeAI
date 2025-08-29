package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Difficulty string

const (
	Easy 		Difficulty = "easy"
	Medium 		Difficulty = "medium"
	Hard 		Difficulty = "hard"
	Nightmare 	Difficulty = "nightmare"
)

type Quiz struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Title 		string				`bson:"title" json:"title"`
	Description string				`bson:"description" json:"description"`
	Difficulty	Difficulty			`bson:"difficulty" json:"difficulty"`
	Quantity	int8				`bson:"quantity" json:"quantity"`
	Retries		int8				`bson:"retries" json:"retries"`
	DocumentID	primitive.ObjectID	`bson:"documentId,omitempty" json:"documentId,omitempty"` //FK -> Document
	CreatedBy 	primitive.ObjectID	`bson:"createdby,omitempty" json:"createdby,omitempty"` // FK -> User
	CreatedAt	primitive.DateTime	`bson:"createdAt" json:"createdAt"`
	Questions 	[]Question			`bson:"questions" json:"questions"`
}

type Question struct{
	Question 	string		`bson:"question" json:"question"`
	Options 	[]string	`bson:"options" json:"options"`
	Answer 		string		`bson:"answer" json:"answer"`
}