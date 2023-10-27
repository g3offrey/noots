package main

import (
	"fmt"
	"net/http"
	"noots/internal/note"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"

	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()

	r.Use(
		render.SetContentType(render.ContentTypeHTML),
		middleware.Logger,
	)

	noteHandlers := NoteHandlers{
		noteService: note.NewService(note.NewInMemoryRepository()),
	}

	r.Get("/", noteHandlers.List())
	r.Get("/notes/{id}", noteHandlers.Single())
	r.Get("/notes/new", noteHandlers.AddForm())
	r.Post("/notes/new", noteHandlers.Add())
	r.Post("/notes/{id}/archive", noteHandlers.Archive())
	r.Post("/notes/{id}/unarchive", noteHandlers.Unarchive())
	r.Post("/notes/search", noteHandlers.Search())
	r.Delete("/notes/{id}", noteHandlers.Delete())

	fmt.Println("Server running on port 3000")
	http.ListenAndServe(":3000", r)
}
