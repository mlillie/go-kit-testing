package account

import "context"

// Services
type Service interface {
	CreatePerson(ctx context.Context, name string) (string, error)
	GetPerson(ctx context.Context, id string) (string, error)
}
