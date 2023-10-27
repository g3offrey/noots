package main

import (
	"errors"
	"fmt"
	"html/template"
	"net/http"
	"noots/internal/note"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

func IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html"))
		t.Execute(w, nil)
	}
}

type NoteHandlers struct {
	noteService *note.Service
}

func (h *NoteHandlers) List() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/list.html"))

		n, err := h.noteService.GetAll()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, n)
	}
}

func (h *NoteHandlers) Single() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/single.html"))

		idFromURL := chi.URLParam(r, "id")
		id := uuid.MustParse(idFromURL)
		n, err := h.noteService.Get(id)
		if err != nil && errors.Is(err, note.NotFoundErr) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.Execute(w, n)
	}
}

func (h *NoteHandlers) AddForm() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/add.html"))

		t.Execute(w, nil)
	}
}

func (h *NoteHandlers) Add() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		n := note.Note{
			Title:     r.FormValue("title"),
			Content:   r.FormValue("content"),
			CreatedAt: time.Now(),
		}
		n.Identify()

		err := h.noteService.Add(n)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		redirectURL := fmt.Sprintf("/notes/%s", n.ID.String())
		w.Header().Add("HX-Redirect", redirectURL)
		w.WriteHeader(http.StatusNoContent)
	}
}

func (h *NoteHandlers) Archive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/single.html"))

		idFromURL := chi.URLParam(r, "id")
		id := uuid.MustParse(idFromURL)

		n, err := h.noteService.Archive(id)
		if err != nil && errors.Is(err, note.NotFoundErr) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.ExecuteTemplate(w, "details", n)
	}
}

func (h *NoteHandlers) Unarchive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/single.html"))

		idFromURL := chi.URLParam(r, "id")
		id := uuid.MustParse(idFromURL)

		n, err := h.noteService.Unarchive(id)
		if err != nil && errors.Is(err, note.NotFoundErr) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.ExecuteTemplate(w, "details", n)
	}
}

func (h *NoteHandlers) Search() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/layout.html", "templates/notes/list.html"))

		search := r.FormValue("search")
		n, err := h.noteService.Search(search)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		t.ExecuteTemplate(w, "notes", n)
	}
}

func (h *NoteHandlers) Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idFromURL := chi.URLParam(r, "id")
		id := uuid.MustParse(idFromURL)

		err := h.noteService.Delete(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Add("HX-Redirect", "/")
		w.WriteHeader(http.StatusNoContent)
	}
}
