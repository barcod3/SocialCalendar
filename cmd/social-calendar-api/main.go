package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/barcod3/socialcalendar/internal/client"
	"github.com/barcod3/socialcalendar/internal/handler"
	"github.com/barcod3/socialcalendar/internal/processor"
	"github.com/barcod3/socialcalendar/internal/router"
	"github.com/kelseyhightower/envconfig"
	"github.com/patrickmn/go-cache"
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
)

type config struct {
	ID        string
	Secret    string
	Username  string
	Password  string
	Subreddit string `default:"londonsocialclub"`
	BaseURL   string `default:"/api/v1"`
	Port      int    `default:"8080"`
}

func main() {

	ctx := context.Background()

	var cfg config

	err := envconfig.Process("SOCIAL_CALENDAR", &cfg)
	if err != nil {
		logrus.WithError(err).Error("Failed to parse config")
		return
	}

	cl, err := client.NewRedditClient(cfg.ID, cfg.Secret, cfg.Username, cfg.Password, cfg.Subreddit)
	if err != nil {
		logrus.WithError(err).Error("Failed to create new reddit client")
		return
	}

	ch := cache.New(15*time.Minute, 30*time.Minute)

	pr := processor.NewEventsProcessor(cl)

	job := func() {
		events, err := pr.ProcessEvents(ctx)
		if err != nil {
			logrus.WithError(err).Error("Failed to process events")
			return
		}
		ch.Set("events", events, cache.DefaultExpiration)
	}

	job()
	cr := cron.New()
	if _, err := cr.AddFunc("@every 10m", job); err != nil {
		logrus.WithError(err).Error("Failed to create cron")
		return
	}
	cr.Start()

	h := handler.NewEventHandler(ch)

	r := router.NewSocialAPI(ctx, cfg.BaseURL, &h)

	srv := http.Server{
		Addr:         fmt.Sprintf(":%v", cfg.Port),
		Handler:      r,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
	}

	logrus.Info("http server started")

	logrus.Fatal(srv.ListenAndServe())

}
