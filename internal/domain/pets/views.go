package pets

import (
	"context"

	"github.com/go-po/po/streams"
)

type Pet struct {
	Id   int64
	Tag  string
	Name string
}

func (pet *Pet) Handle(ctx context.Context, msg streams.Message) error {
	switch event := msg.Data.(type) {
	case Added:
		pet.Name = event.Name
	default:
	}
	return nil
}
