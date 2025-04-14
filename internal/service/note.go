package service

import "github.com/rtzgod/simple-crud/internal/models"

type Repository interface {
	CreateNote(title, content string) (id int, err error)
	GetNotes() ([]models.Note, error)
	UpdateNote(id int, title, content string) error
	DeleteNote(id int) error
}

type Service struct {
	repo Repository
}

func NewService(repo Repository) *Service {
	return &Service{
		repo: repo,
	}
}

func (s *Service) CreateNote(title, content string) (id int, err error) {
	id, err = s.repo.CreateNote(title, content)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (s *Service) GetNotes() ([]models.Note, error) {
	notes, err := s.repo.GetNotes()
	if err != nil {
		return []models.Note{}, err
	}
	return notes, nil
}
func (s *Service) UpdateNote(id int, title, content string) error {
	err := s.repo.UpdateNote(id, title, content)
	if err != nil {
		return err
	}
	return nil
}
func (s *Service) DeleteNote(id int) error {
	err := s.repo.DeleteNote(id)
	if err != nil {
		return err
	}
	return nil
}
