package service

import (
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (string, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, string, error)
}

type Note interface {
	Create(userId int, note model.Note) (int, error)
	GetAll(userId int) ([]model.Note, error)
	GetById(userId, noteId int) (model.Note, error)
	Delete(userId, noteId int) error
	Update(userId int, noteId int, input model.UpdateNoteInput) error
}

type Service struct {
	Authorization
	Note
}

func NewService(repos *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repos.Authorization),
		Note:          NewNoteService(repos.Note),
	}
}
