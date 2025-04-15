package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/rtzgod/simple-crud/internal/models"
)

// Success response structure
type successResponse struct {
	ID int `json:"id"`
}

func (h *Handler) createNote(w http.ResponseWriter, r *http.Request) {
	var note models.Note

	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	if note.Title == "" || note.Content == "" {
		http.Error(w, "Title and Content are required", http.StatusBadRequest)
		return
	}

	id, err := h.service.CreateNote(note.Title, note.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(successResponse{ID: id})

}

func (h *Handler) getNotes(w http.ResponseWriter, r *http.Request) {
	notes, err := h.service.GetNotes()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(notes); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
		return
	}
}

func (h *Handler) updateNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	var note models.Note
	if err := json.NewDecoder(r.Body).Decode(&note); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	err = h.service.UpdateNote(noteID, note.Title, note.Content)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *Handler) deleteNote(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	noteID, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, "Invalid note ID", http.StatusBadRequest)
		return
	}

	err = h.service.DeleteNote(noteID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
