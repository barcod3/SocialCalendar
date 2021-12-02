// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/barcod3/socialcalendar/internal/processor (interfaces: RedditClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	reddit "github.com/vartanbeno/go-reddit/v2/reddit"
)

// MockRedditClient is a mock of RedditClient interface.
type MockRedditClient struct {
	ctrl     *gomock.Controller
	recorder *MockRedditClientMockRecorder
}

// MockRedditClientMockRecorder is the mock recorder for MockRedditClient.
type MockRedditClientMockRecorder struct {
	mock *MockRedditClient
}

// NewMockRedditClient creates a new mock instance.
func NewMockRedditClient(ctrl *gomock.Controller) *MockRedditClient {
	mock := &MockRedditClient{ctrl: ctrl}
	mock.recorder = &MockRedditClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRedditClient) EXPECT() *MockRedditClientMockRecorder {
	return m.recorder
}

// NewPosts mocks base method.
func (m *MockRedditClient) NewPosts(arg0 context.Context) ([]*reddit.Post, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewPosts", arg0)
	ret0, _ := ret[0].([]*reddit.Post)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewPosts indicates an expected call of NewPosts.
func (mr *MockRedditClientMockRecorder) NewPosts(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewPosts", reflect.TypeOf((*MockRedditClient)(nil).NewPosts), arg0)
}