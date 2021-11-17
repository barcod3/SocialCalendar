package main

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/barcod3/socialcalendar/internal/client"
	"github.com/barcod3/socialcalendar/internal/handler"
	"github.com/barcod3/socialcalendar/internal/router"
	"github.com/kelseyhightower/envconfig"
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

	c, err := client.NewRedditClient(cfg.ID, cfg.Secret, cfg.Username, cfg.Password, cfg.Subreddit)
	if err != nil {
		logrus.WithError(err).Error("Failed to create new reddit client")
		return
	}

	h := handler.NewEventHandler(c)

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
