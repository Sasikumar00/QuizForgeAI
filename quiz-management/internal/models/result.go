package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Result struct {
	ID 			primitive.ObjectID 		`bson:"_id,omitempty" json:"id"`
	UserID 		primitive.ObjectID 		`bson:"userId,omitempty" json:"userId,omitempty"` // FK -> User
	QuizId 		primitive.ObjectID 		`bson:"quizId,omitempty" json:"quizId,omitempty"` // FK -> Quiz
	Score 		int						`bson:"score" json:"score"`
	SubmittedAt int64 					`bson:"submittedAt" json:"submittedAt"`
}