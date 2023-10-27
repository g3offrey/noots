package note

import (
	"errors"

	"github.com/google/uuid"
)

type Service struct {
	repository Repository
}

var NotFoundErr = errors.New("note not found")

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetAll() ([]Note, error) {
	return s.repository.FindAll()
}

func (s *Service) Get(id uuid.UUID) (note Note, err error) {
	note, found, err := s.repository.FindByID(id)
	if err != nil {
		return
	}

	if !found {
		err = NotFoundErr
		return
	}

	return
}

func (s *Service) Add(note Note) error {
	return s.repository.Add(note)
}

func (s *Service) Archive(id uuid.UUID) (note Note, err error) {
	note, found, err := s.repository.FindByID(id)
	if err != nil {
		return
	}

	if !found {
		return
	}

	note.Archived = true

	err = s.repository.Update(id, note)
	if err != nil {
		return
	}

	return
}

func (s *Service) Unarchive(id uuid.UUID) (note Note, err error) {
	note, found, err := s.repository.FindByID(id)
	if err != nil {
		return
	}

	if !found {
		return
	}

	note.Archived = false

	err = s.repository.Update(id, note)
	if err != nil {
		return
	}

	return
}

func (s *Service) Search(search string) (matches []Note, err error) {
	return s.repository.FindByTitle(search)
}

func (s *Service) Delete(id uuid.UUID) error {
	return s.repository.Delete(id)
}
