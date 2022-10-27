package repository

import (
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GetUser(username string) (model.User, error)
}

type Note interface {
	Create(userId int, note model.Note) (int, error)
	GetAll(userId int) ([]model.Note, error)
	GetById(userId, noteId int) (model.Note, error)
	Delete(userId, noteId int) error
	Update(userId int, noteId int, input model.UpdateNoteInput) error
}

type Repository struct {
	Authorization
	Note
}

func NewRepository(db *mongo.Database) *Repository {
	return &Repository{
		Authorization: NewAuthMongo(db),
		Note:          NewNoteMongo(db),
	}
}
