package pets

import (
	"context"
	"github.com/go-po/po"
	"github.com/go-po/po/stream"
)

type AddPetHandler struct {
	es EventStore
}

func (h *AddPetHandler) Handle(ctx context.Context, msg stream.Message) error {
	streamId := StreamPets.WithEntity("%d", msg.Number)
	switch cmd := msg.Data.(type) {
	case AddPetCmd:
		return addPet(h.es.Stream(ctx, streamId.String()), cmd)
	default:
		// ignore
	}
	return nil
}

func addPet(petStream *po.Stream, cmd AddPetCmd) error {
	tx, err := petStream.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	pos, err := petStream.Size()
	if err != nil {
		return err
	}
	if pos > 0 {
		// already handled
		return nil
	}
	petStream.AppendTx(tx, Added{
		Name: cmd.Name,
		Tags: cmd.Tags,
	})
	_, err = tx.Commit()
	return err
}
