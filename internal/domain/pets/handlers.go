package pets

import (
	"context"

	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

type AddPetHandler struct {
	es EventStore
}

func (h *AddPetHandler) Handle(ctx context.Context, msg streams.Message) error {
	streamId := StreamPets.WithEntity("%d", msg.Number)
	switch cmd := msg.Data.(type) {
	case AddPetCmd:
		return addPet(h.es.Stream(ctx, streamId), cmd)
	default:
		// ignore
	}
	return nil
}

func addPet(petStream *po.Stream, cmd AddPetCmd) error {
	pos, err := petStream.Size()
	if err != nil {
		return err
	}
	if pos > 0 {
		// already handled
		return nil
	}
	_, err = petStream.Append(Added{
		Name: cmd.Name,
		Tags: cmd.Tags,
	})
	return err
}
