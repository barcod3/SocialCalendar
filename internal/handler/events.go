package handler

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/barcod3/socialcalendar/internal/model"
	"github.com/sirupsen/logrus"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type EventHandler struct {
	client RedditClient
}

func NewEventHandler(client RedditClient) EventHandler {
	return EventHandler{
		client: client,
	}
}

type RedditClient interface {
	NewPosts(ctx context.Context) ([]*reddit.Post, error)
}

func (h *EventHandler) GetEvents(ctx context.Context) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		posts, err := h.client.NewPosts(ctx)
		if err != nil {
			JSONError(w, ErrorMessage{
				Message: "failed to get posts",
				Error:   err,
				Code:    http.StatusBadGateway,
			})
			return
		}

		out := []model.Event{}

		for _, post := range posts {
			event, err := model.ParseEvent(post)
			if err != nil {
				if err == model.ErrNotEvent {
					continue
				}
				JSONError(w, ErrorMessage{
					Message: "failed to parse posts",
					Error:   err,
					Code:    http.StatusInternalServerError,
				})
				return
			}
			out = append(out, *event)
		}

		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(out); err != nil {
			logrus.WithError(err).Error("Couldnt output json events")
		}

	})
}
