package processor_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/barcod3/socialcalendar/internal/processor"
	"github.com/barcod3/socialcalendar/internal/processor/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

//go:generate mockgen -destination=./mocks/mock_RedditClient.go -package=mocks github.com/barcod3/socialcalendar/internal/processor RedditClient

func TestProcessEvents(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mockClient := mocks.NewMockRedditClient(ctrl)
	ctx := context.Background()

	now := time.Now()

	posts := []*reddit.Post{
		{
			ID:    "Test1",
			Title: "Not an event",
		},
		{
			ID:    "Test2",
			Title: fmt.Sprintf("[%s] four digit event", now.Format("02/01/2006")),
		},
		{
			ID:    "Test3",
			Title: fmt.Sprintf("[%s] two digit event", now.Format("02/01/06")),
		},
		{
			ID:    "Test4",
			Title: fmt.Sprintf("[%s] future event (ignored)", now.Add(time.Hour*48).Format("02/01/2006")),
		},
		{
			ID:    "Test5",
			Title: fmt.Sprintf("[%s] past event (ignored)", now.Add(time.Hour*-48).Format("02/01/2006")),
		},
	}

	p := processor.NewEventsProcessor(mockClient)

	mockClient.EXPECT().NewPosts(ctx).Return(posts, nil)

	events, err := p.ProcessEvents(ctx)

	assert.NoError(t, err, "Unexpected error calling ProcessEvents")
	assert.Len(t, events, 3)

}
