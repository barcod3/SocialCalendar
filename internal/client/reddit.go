package client

import (
	"context"

	"github.com/pkg/errors"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type RedditClient struct {
	id, secret, username, password, subreddit string
	client                                    *reddit.Client
}

func NewRedditClient(id, secret, username, password, subreddit string) (*RedditClient, error) {
	rc := RedditClient{
		id:        id,
		secret:    secret,
		username:  username,
		password:  password,
		subreddit: subreddit,
	}
	credentials := reddit.Credentials{ID: id, Secret: secret, Username: username, Password: password}
	client, err := reddit.NewClient(credentials)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to create new client")
	}
	rc.client = client
	return &rc, nil
}

func (rc RedditClient) NewPosts(ctx context.Context) ([]*reddit.Post, error) {
	posts, _, err := rc.client.Subreddit.NewPosts(ctx, "londonsocialclub", &reddit.ListOptions{
		Limit: 100,
	})
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get new posts")
	}
	return posts, nil
}
