package router

import (
	"context"
	"net/http"

	"github.com/gorilla/mux"
)

type EventHandler interface {
	GetEvents(ctx context.Context) http.HandlerFunc
}

func NewSocialAPI(ctx context.Context, baseurl string, handler EventHandler) *mux.Router {
	r := mux.NewRouter()
	sr := r.PathPrefix(baseurl).Subrouter()
	sr.HandleFunc("/events", handler.GetEvents(ctx))
	return r
}
