package repository

import (
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/rtzgod/simple-crud/internal/models"
)

type Repository struct {
	db *sqlx.DB
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) CreateNote(title, content string) (id int, err error) {
	stmt, err := r.db.Prepare("INSERT INTO notes (title, content) VALUES ($1, $2) RETURNING id")
	if err != nil {
		return 0, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(title, content).Scan(&id)
	if err != nil {
		return 0, err
	}
	return id, nil
}
func (r *Repository) GetNotes() ([]models.Note, error) {
	query := "SELECT id, title, content FROM notes"

	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var notes []models.Note

	for rows.Next() {
		var note models.Note
		if err := rows.Scan(&note.ID, &note.Title, &note.Content); err != nil {
			return nil, err
		}
		notes = append(notes, note)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return notes, nil
}
func (r *Repository) UpdateNote(id int, title, content string) error {
	stmt, err := r.db.Prepare("UPDATE notes SET title = $1, content = $2 WHERE id = $3")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(title, content, id)
	if err != nil {
		return err
	}

	return nil
}

func (r *Repository) DeleteNote(id int) error {
	stmt, err := r.db.Prepare("DELETE FROM notes WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(id)
	if err != nil {
		return err
	}

	// Проверяем количество затронутых строк
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// Если ни одна строка не была удалена, возвращаем ошибку
	if rowsAffected == 0 {
		return errors.New("заметка не найдена")
	}

	return nil
}
