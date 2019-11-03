package account

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

type (
	CreatePersonRequest struct {
		Name string `json:"name"`
	}

	CreatePersonResponse struct {
		Ok string `json: "ok"`
	}

	GetPersonRequest struct {
		Id string `json:"id"`
	}

	GetPersonResponse struct {
		Name string `json:"name"`
	}
)

func encodeResponse(ctx context.Context, writer http.ResponseWriter, response interface{}) error {
	return json.NewEncoder(writer).Encode(response)
}

func decodePersonReq(ctx context.Context, request *http.Request) (interface{}, error) {
	var req CreatePersonRequest
	err := json.NewDecoder(request.Body).Decode(&req)

	if err != nil {
		return nil, err
	}

	return req, nil
}

func decodeNameReq(ctx context.Context, request *http.Request) (interface{}, error) {
	var req GetPersonRequest
	vars := mux.Vars(request)

	req = GetPersonRequest{
		Id: vars["id"],
	}

	return req, nil
}
