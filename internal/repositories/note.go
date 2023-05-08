package repositories

import (
	"context"
	"errors"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"

	"github.com/Yu-Leo/vk-internship-tarantool-2023/internal/models"
	"github.com/Yu-Leo/vk-internship-tarantool-2023/pkg/postgresql"
)

type NoteRepository interface {
	Set(userID string, note *models.Note) error
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

func (r *noteRepository) Set(userID string, note *models.Note) (err error) {
	var pgErr *pgconn.PgError

	q1 := `
INSERT INTO notes (user_id, service, login, password)
VALUES ($1, $2, $3, $4);`
	_, err = r.postgresConnection.Exec(context.Background(), q1,
		userID, note.ServiceName, note.Login, note.Password)
	if err == nil {
		return nil
	}

	if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
		q2 := `
UPDATE notes
SET login = $1, password = $2
WHERE user_id = $3 AND service = $4;`

		_, err = r.postgresConnection.Exec(context.Background(), q2,
			note.Login, note.Password, userID, note.ServiceName)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *noteRepository) Get(userID, serviceName string) (_ *models.Note, err error) {
	q := `
SELECT (service, login, password)
FROM notes
WHERE notes.user_id = $1 AND service = $2;`

	var note models.Note
	err = r.postgresConnection.QueryRow(context.Background(), q, userID, serviceName).
		Scan(&note.ServiceName, &note.Login, &note.Password)
	if err != nil {
		return nil, err
	}
	return &note, nil
}

func (r *noteRepository) Del(userID, serviceName string) (err error) {
	q := `
DELETE
FROM notes
WHERE user_id = $1 AND service = $2;`
	_, err = r.postgresConnection.Exec(context.Background(), q, userID, serviceName)
	if err != nil {
		return err
	}
	return nil
}
