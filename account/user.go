package account

import "context"

type Person struct {
	ID   string `json:"id,omitempty"`
	Name string `json:"name"`
}

type Repository interface {
	CreatePerson(ctx context.Context, person Person) error
	GetPerson(ctx context.Context, id string) (string, error)
}
