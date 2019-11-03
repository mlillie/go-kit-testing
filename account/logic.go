package account

import (
	"context"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/gofrs/uuid"
)

type service struct {
	repostiry Repository
	logger    log.Logger
}

func NewService(rep Repository, logger log.Logger) Service {
	return &service{
		repostiry: rep,
		logger:    logger,
	}
}

func (s service) CreatePerson(ctx context.Context, name string) (string, error) {
	logger := log.With(s.logger, "method", "CreatePerson")

	uuid, _ := uuid.NewV4()
	id := uuid.String()
	person := Person{
		ID:   id,
		Name: name,
	}

	if err := s.repostiry.CreatePerson(ctx, person); err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("create person", id)

	return "Success", nil
}

func (s service) GetPerson(ctx context.Context, id string) (string, error) {
	logger := log.With(s.logger, "method", "GetPerson")

	name, err := s.repostiry.GetPerson(ctx, id)

	if err != nil {
		level.Error(logger).Log("err", err)
		return "", err
	}

	logger.Log("Get person", id)

	return name, nil
}
