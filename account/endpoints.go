package account

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

type Endpoints struct {
	CreatePerson endpoint.Endpoint
	GetPerson    endpoint.Endpoint
}

func MakeEndpoints(s Service) Endpoints {
	return Endpoints{
		CreatePerson: makeCreatePersonEndpoint(s),
		GetPerson:    makeGetPersonEndpoint(s),
	}
}

func makeCreatePersonEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(CreatePersonRequest)
		ok, err := s.CreatePerson(ctx, req.Name)
		return CreatePersonResponse{Ok: ok}, err
	}
}

func makeGetPersonEndpoint(s Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(GetPersonRequest)
		name, err := s.GetPerson(ctx, req.Id)
		return GetPersonResponse{
			Name: name,
		}, err
	}
}
