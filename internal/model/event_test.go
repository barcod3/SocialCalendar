package model_test

import (
	"reflect"
	"testing"
	"time"

	"github.com/barcod3/socialcalendar/internal/model"
	"github.com/vartanbeno/go-reddit/v2/reddit"
)

func TestParseEvent(t *testing.T) {
	tests := []struct {
		name    string
		arg     reddit.Post
		want    *model.Event
		wantErr bool
	}{
		{
			name: "test no parenthesis",
			arg: reddit.Post{
				Title: "a reddit post",
			},
			wantErr: true,
		},
		{
			name: "test not date",
			arg: reddit.Post{
				Title: "[NOT DATE] a reddit post",
			},
			wantErr: true,
		},
		{
			name: "test two digit year",
			arg: reddit.Post{
				Title: "[01/01/01] a reddit post",
			},
			wantErr: false,
			want: &model.Event{
				Title: "a reddit post",
				Date:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		},
		{
			name: "test four digit year",
			arg: reddit.Post{
				Title: "[01/01/2001] a reddit post",
			},
			wantErr: false,
			want: &model.Event{
				Title: "a reddit post",
				Date:  time.Date(2001, 1, 1, 0, 0, 0, 0, time.UTC),
			},
		}, {
			name: "test not US date",
			arg: reddit.Post{
				Title: "[13/12/2001] a reddit post",
			},
			wantErr: false,
			want: &model.Event{
				Title: "a reddit post",
				Date:  time.Date(2001, 12, 13, 0, 0, 0, 0, time.UTC),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := model.ParseEvent(&tt.arg)
			if (err != nil) != tt.wantErr {
				t.Errorf("ParseEvent() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseEvent() = %v, want %v", got, tt.want)
			}
		})
	}
}
