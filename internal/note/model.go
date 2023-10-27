package note

import (
	"github.com/google/uuid"
	"time"
)

type Note struct {
	ID        uuid.UUID
	Title     string
	Content   string
	Archived  bool
	CreatedAt time.Time
}

func (n *Note) Identify() {
	n.ID = uuid.New()
}
