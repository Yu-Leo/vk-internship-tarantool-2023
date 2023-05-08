package repositories

import (
	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/models"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/pkg/postgresql"
)

type NoteRepository interface {
	Set(userID string, note models.Note) error
	Get(userID, serviceName string) (*models.Note, error)
	Del(userID, serviceName string) error
}

type noteRepository struct {
	postgresConnection postgresql.Connection
}

func NewPostgresNoteRepository(pc postgresql.Connection) NoteRepository {
	return &noteRepository{
		postgresConnection: pc,
	}
}

func (r *noteRepository) Set(userID string, note models.Note) error {
	return nil
}

func (r *noteRepository) Get(userID, serviceName string) (*models.Note, error) {
	return nil, nil
}

func (r *noteRepository) Del(userID, serviceName string) error {
	return nil
}
