package pets

import (
	"context"
	"github.com/go-po/po/stream"
)

type Pet struct {
	Id   int64
	Tag  string
	Name string
}

func (pet *Pet) Handle(ctx context.Context, msg stream.Message) error {
	switch event := msg.Data.(type) {
	case Added:
		pet.Name = event.Name
	default:
	}
	return nil
}

type DeletedView struct {
	Deleted bool
}

func (view *DeletedView) Handle(ctx context.Context, msg stream.Message) error {
	switch msg.Data.(type) {
	case Deleted:
		view.Deleted = true
	default:
	}
	return nil
}
