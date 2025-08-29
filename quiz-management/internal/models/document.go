package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Document struct {
	ID 			primitive.ObjectID	`bson:"_id,omitempty" json:"id"`
	Name		string				`bson:"name" json:"name"`
	Reuploads	int					`bson:"reuploads" json:"reuploads"`
	UploadedBy 	primitive.ObjectID 	`bson:"uploadedBy,omitempty" json:"uploadedBy,omitempty"` // FK -> User
	UploadedAt 	int64				`bson:"uploadedAt" json:"uploadedAt"`
}