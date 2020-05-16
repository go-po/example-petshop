package pets

import (
	"context"

	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

type DeleteAction struct {
	Deleted bool
}

func (view *DeleteAction) Execute(appender po.TransactionAppender) error {
	if appender.Size() == 0 || view.Deleted {
		// nothing to do
		return nil
	}
	appender.Append(Deleted{})
	return nil
}

func (view *DeleteAction) Handle(ctx context.Context, msg streams.Message) error {
	switch msg.Data.(type) {
	case Deleted:
		view.Deleted = true
	default:
	}
	return nil
}
