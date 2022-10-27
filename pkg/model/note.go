package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct {
	ID      primitive.ObjectID `bson:"_id"`
	Id      int                `json:"id"`
	Title   string             `json:"title" binding:"required"`
	Content string             `json:"content"`
}

type UsersNote struct {
	Id     int
	UserId int
	NoteId int
}

type UsersNotes struct {
	Id     primitive.ObjectID `bson:"_id"`
	UserId int
	NoteId int
}

type UpdateNoteInput struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
