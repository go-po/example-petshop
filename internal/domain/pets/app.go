package pets

import (
	"context"

	"github.com/go-po/po"
	"github.com/go-po/po/streams"
)

var (
	AddsStream = streams.ParseId("pets:add")
	StreamPets = streams.ParseId("pets")
)

type EventStore interface {
	Stream(ctx context.Context, streamId streams.Id) *po.Stream
	Project(ctx context.Context, streamId streams.Id, projection streams.Handler) error
	Subscribe(ctx context.Context, subscriptionId string, streamId streams.Id, subscriber interface{}) error
}

func New(es EventStore) (*App, error) {

	err := es.Subscribe(context.Background(), "add-pets-handler-1", AddsStream, &AddPetHandler{
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
	return app.es.
		Stream(ctx, streams.ParseId("pets-%d", id)).
		Execute(&DeleteAction{})
}

func (app *App) GetPetById(ctx context.Context, id int64) (Pet, error) {
	pet := Pet{
		Id: id,
	}
	err := app.es.Project(ctx, streams.ParseId("pets-%d", id), &pet)
	if err != nil {
		return Pet{}, err
	}
	return pet, nil
}

func (app *App) AddPet(ctx context.Context, name string, tags []string) (int64, error) {
	// use the message numbering on the 'pets:add'
	// stream as a sequence for Pet ids
	id, err := app.es.
		Stream(ctx, AddsStream).
		Append(AddPetCmd{
			Name: name,
			Tags: tags,
		})
	if err != nil {
		return 0, err
	}
	return id, nil
}
