package users

import (
	"context"
	"github.com/go-po/po"
	"github.com/go-po/po/stream"
)

type EventStore interface {
	Stream(ctx context.Context, streamId string) *po.Stream
	Project(ctx context.Context, streamId string, projection stream.Handler) error
	Subscribe(ctx context.Context, subscriptionId, streamId string, subscriber interface{}) error
}

func New(es EventStore) (*App, error) {
	return &App{es: es}, nil
}

type App struct {
	es EventStore
}
