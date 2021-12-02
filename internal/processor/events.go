package processor

import (
	"context"
	"sort"
	"time"

	"github.com/barcod3/socialcalendar/internal/model"
	"github.com/pkg/errors"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type EventsProcessor struct {
	client RedditClient
}

type RedditClient interface {
	NewPosts(ctx context.Context) ([]*reddit.Post, error)
}

func NewEventsProcessor(client RedditClient) EventsProcessor {
	return EventsProcessor{
		client: client,
	}
}

func (e EventsProcessor) ProcessEvents(ctx context.Context) ([]*model.Event, error) {
	posts, err := e.client.NewPosts(ctx)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get posts")
	}

	out := []*model.Event{}

	now := time.Now().Truncate(time.Hour * 24)

	for _, post := range posts {
		event, err := model.ParseEvent(post)
		if err != nil {
			if err == model.ErrNotEvent {
				continue
			}
			return nil, errors.Wrap(err, "failed to parse posts")
		}

		if event.Date.After(now) || event.Date.Equal(now) {
			out = append(out, event)
		}
	}

	sort.SliceStable(out, func(i, j int) bool {
		return out[i].Date.Before(out[j].Date)
	})

	return out, nil

}
