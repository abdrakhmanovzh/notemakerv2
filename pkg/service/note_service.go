package service

import (
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/model"
	"github.com/abdrakhmanovzh/notemaker2.0/pkg/repository"
)

type NoteService struct {
	repo repository.Note
}

func NewNoteService(repo repository.Note) *NoteService {
	return &NoteService{repo: repo}
}

func (s *NoteService) Create(userId int, note model.Note) (int, error) {
	return s.repo.Create(userId, note)
}

func (s *NoteService) GetAll(userId int) ([]model.Note, error) {
	return s.repo.GetAll(userId)
}

func (s *NoteService) GetById(userId, noteId int) (model.Note, error) {
	return s.repo.GetById(userId, noteId)
}

func (s *NoteService) Delete(userId, noteId int) error {
	return s.repo.Delete(userId, noteId)
}

func (s *NoteService) Update(userId int, noteId int, input model.UpdateNoteInput) error {
	return s.repo.Update(userId, noteId, input)
}
