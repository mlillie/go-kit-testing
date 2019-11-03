package account

import (
	"context"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
)

func NewHTTPServer(ctx context.Context, endpoints Endpoints) http.Handler {
	router := mux.NewRouter()
	router.Use(commonMiddleware)

	router.Methods("POST").Path("/person").Handler(httptransport.NewServer(
		endpoints.CreatePerson,
		decodePersonReq,
		encodeResponse,
	))

	router.Methods("GET").Path("/person/{id}").Handler(httptransport.NewServer(
		endpoints.GetPerson,
		decodeNameReq,
		encodeResponse,
	))

	return router
}

func commonMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		writer.Header().Add("Content-Type", "application/json")
		next.ServeHTTP(writer, request)
	})
}
