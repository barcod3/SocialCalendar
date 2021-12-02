package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/patrickmn/go-cache"
	"github.com/sirupsen/logrus"
)

type EventHandler struct {
	cache *cache.Cache
}

func NewEventHandler(cache *cache.Cache) EventHandler {
	return EventHandler{
		cache: cache,
	}
}

func (h *EventHandler) GetEvents(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		events, found := h.cache.Get("events")
		if !found {
			JSONError(w, ErrorMessage{
				Message: "events cache empty",
				Error:   nil,
				Code:    500,
			})
			return
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(events); err != nil {
			logrus.WithError(err).Error("Couldnt output json events")
		}

	})
}
