package handler

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rtzgod/simple-crud/internal/models"
)

type Service interface {
	CreateNote(title, content string) (id int, err error)
	GetNotes() ([]models.Note, error)
	UpdateNote(id int, title, content string) error
	DeleteNote(id int) error
}

type Handler struct {
	service Service
}

func NewHandler(service Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.Logger)

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("OK!"))
	})
	r.Mount("/", h.Routes())

	return r
}

func (h *Handler) Routes() chi.Router {
	r := chi.NewRouter()

	r.Post("/notes", h.createNote)
	r.Get("/notes", h.getNotes)
	r.Put("/notes/{id}", h.updateNote)
	r.Delete("/notes/{id}", h.deleteNote)

	return r
}
