package account

import (
	"context"
	"database/sql"
	"errors"

	"github.com/go-kit/kit/log"
)

var RepoErr = errors.New("Unable to handle repo request")

type repo struct {
	db     *sql.DB
	logger log.Logger
}

func NewRepo(db *sql.DB, logger log.Logger) Repository {
	return &repo{
		db:     db,
		logger: log.With(logger, "repo", "sql"),
	}
}

func (repo *repo) CreatePerson(ctx context.Context, person Person) error {
	sqlStatement := `INSERT INTO person (id, name) VALUES ($1, $2)`

	if person.Name == "" {
		return RepoErr
	}

	_, err := repo.db.ExecContext(ctx, sqlStatement, person.ID, person.Name)

	if err != nil {
		return err
	}

	return nil
}

func (repo *repo) GetPerson(ctx context.Context, id string) (string, error) {
	var name string
	err := repo.db.QueryRow("SELECT name FROM person WHERE id=$1", id).Scan(&name)

	if err != nil {
		return "", RepoErr
	}

	return name, nil
}
