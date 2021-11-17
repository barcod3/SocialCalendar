package model

import (
	"strings"
	"time"

	"github.com/pkg/errors"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

type Event struct {
	ID        string
	Title     string
	Body      string
	Date      time.Time
	Author    string
	VoteCount int
	Comments  int
}

var ErrNotEvent = errors.New("Is not an event")

func ParseEvent(p *reddit.Post) (*Event, error) {
	if !strings.HasPrefix(p.Title, "[") {
		return nil, ErrNotEvent
	}

	s := strings.Split(p.Title, "]")

	ds := strings.Replace(s[0], "[", "", 1)

	var d time.Time

	d, err := time.Parse("02/01/2006", ds)
	if err != nil {
		d, err = time.Parse("02/01/06", ds)
		if err != nil {
			return nil, ErrNotEvent
		}
	}

	return &Event{
		ID:        p.ID,
		Title:     strings.Trim(s[1]," "),
		Body:      p.Body,
		Date:      d,
		Author:    p.Author,
		VoteCount: p.Score,
		Comments:  p.NumberOfComments,
	}, nil
}
