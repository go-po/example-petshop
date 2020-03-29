package pets

import (
	"context"
	"github.com/go-po/po"
	"github.com/go-po/po/stream"
)

var (
	StreamPetsAdd = "pets:add"
	StreamPets    = stream.ParseId("pets")
)

type EventStore interface {
	Stream(ctx context.Context, streamId string) *po.Stream
	Project(ctx context.Context, streamId string, projection stream.Handler) error
	Subscribe(ctx context.Context, subscriptionId, streamId string, subscriber interface{}) error
}

func New(es EventStore) (*App, error) {

	err := es.Subscribe(context.Background(), "add-pets-handler-1", StreamPetsAdd, &AddPetHandler{
		es: es,
	})
	if err != nil {
		return nil, err
	}

	return &App{es: es}, nil
}

type App struct {
	es EventStore
}

func (app *App) DeletePet(ctx context.Context, id int64) error {
	stream := app.es.
		Stream(ctx, StreamPets.WithEntity("%d", id).String())

	tx, err := stream.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	view := DeletedView{}
	err = stream.Project(&view)
	if err != nil {
		return err
	}
	size, err := stream.Size()
	if err != nil {
		return err
	}

	if view.Deleted || size == 0 {
		// nothing to do
		return nil
	}

	stream.AppendTx(tx, Deleted{})

	_, err = tx.Commit()
	if err != nil {
		return err
	}

	return nil
}

func (app *App) GetPetById(ctx context.Context, id int64) (Pet, error) {
	pet := Pet{
		Id: id,
	}
	err := app.es.Project(ctx, StreamPets.WithEntity("%d", id).String(), &pet)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (app *App) AddPet(ctx context.Context, name string, tags []string) (int64, error) {
	// use the message numbering on the 'pets:add'
	// stream as a sequence for Pet ids
	id, err := app.es.
		Stream(ctx, StreamPetsAdd).
		Append(AddPetCmd{
			Name: name,
			Tags: tags,
		})
	if err != nil {
		return 0, err
	}
	return id, nil
}
