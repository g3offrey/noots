package note

import (
	"fmt"
	"strings"
	"time"

	"github.com/google/uuid"
)

type Repository interface {
	FindAll() ([]Note, error)
	FindByID(id uuid.UUID) (Note, bool, error)
	FindByTitle(title string) ([]Note, error)
	Add(note Note) error
	Update(id uuid.UUID, note Note) error
	Delete(id uuid.UUID) error
}

type InMemoryRepository struct {
	notes []Note
}

func NewInMemoryRepository() *InMemoryRepository {
	r := InMemoryRepository{notes: make([]Note, 10)}

	for i := range r.notes {
		n := Note{
			Title:     fmt.Sprintf("Note %d", i),
			Content:   fmt.Sprintf("This is my note number %d", i),
			CreatedAt: time.Now(),
		}
		n.Identify()

		r.notes[i] = n
	}

	return &r
}

func (r *InMemoryRepository) FindAll() ([]Note, error) {
	return r.notes, nil
}

func (r *InMemoryRepository) FindByID(id uuid.UUID) (n Note, found bool, err error) {
	notes, err := r.FindAll()
	if err != nil {
		return
	}

	for _, n := range notes {
		if n.ID == id {
			return n, true, nil
		}
	}

	return
}

func (r *InMemoryRepository) FindByTitle(title string) (matches []Note, err error) {
	notes, err := r.FindAll()
	if err != nil {
		return
	}

	for _, n := range notes {
		if strings.Contains(
			strings.ToLower(n.Title),
			strings.ToLower(title),
		) {
			matches = append(matches, n)
		}
	}

	return
}

func (r *InMemoryRepository) Add(note Note) error {
	r.notes = append(r.notes, note)

	return nil
}

func (r *InMemoryRepository) Update(id uuid.UUID, note Note) error {
	for i, n := range r.notes {
		if n.ID == id {
			r.notes[i] = note
		}
	}

	return nil
}

func (r *InMemoryRepository) Delete(id uuid.UUID) error {
	for i, n := range r.notes {
		if n.ID == id {
			r.notes = append(r.notes[:i], r.notes[i+1:]...)
			break
		}
	}

	return nil
}
