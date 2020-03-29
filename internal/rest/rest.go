package rest

import (
	"context"
	"encoding/json"
	"github.com/go-po/example-petstore/generated/petstore"
	"github.com/go-po/example-petstore/internal/domain/pets"
	"github.com/labstack/echo/v4"
)

type PetsApp interface {
	AddPet(ctx context.Context, name string, tags []string) (int64, error)
	GetPetById(ctx context.Context, id int64) (pets.Pet, error)
	DeletePet(ctx context.Context, id int64) error
}

func New(petsApp PetsApp) *Api {
	return &Api{pets: petsApp}
}

type Api struct {
	pets PetsApp
}

func (api *Api) FindPets(ctx echo.Context, params petstore.FindPetsParams) error {
	return nil
}

func (api *Api) AddPet(ctx echo.Context) error {
	newPet := petstore.NewPet{}
	err := json.NewDecoder(ctx.Request().Body).Decode(&newPet)
	if err != nil {
		return err
	}
	var tags []string
	if newPet.Tag != nil {
		tags = append(tags, *newPet.Tag)
	}
	id, err := api.pets.AddPet(ctx.Request().Context(), newPet.Name, tags)
	if err != nil {
		return err
	}
	return writePet(ctx.Response(), petstore.Pet{
		NewPet: newPet,
		Id:     id,
	})
}

func (api *Api) DeletePet(ctx echo.Context, id int64) error {
	err := api.pets.DeletePet(ctx.Request().Context(), id)
	if err != nil {
		return err
	}
	return nil
}

func (api *Api) FindPetById(ctx echo.Context, id int64) error {
	pet, err := api.pets.GetPetById(ctx.Request().Context(), id)
	if err != nil {
		return err
	}
	var tag *string
	if len(pet.Tag) > 0 {
		tag = &pet.Tag
	}
	err = writePet(ctx.Response(), petstore.Pet{
		NewPet: petstore.NewPet{
			Name: pet.Name,
			Tag:  tag,
		},
		Id: pet.Id,
	})
	if err != nil {
		return err
	}
	return nil
}

func writePet(response *echo.Response, pet petstore.Pet) error {
	return json.NewEncoder(response).Encode(pet)
}
