package rest

import (
	"github.com/go-po/example-petstore/generated/petstore"
	"github.com/labstack/echo/v4"
)

type PetsApp interface {
}

func New(petsApp PetsApp) *Api {
	return &Api{pets: petsApp}
}

type Api struct {
	pets PetsApp
}

func (api *Api) FindPets(ctx echo.Context, params petstore.FindPetsParams) error {
	panic("implement me")
}

func (api *Api) AddPet(ctx echo.Context) error {
	panic("implement me")
}

func (api *Api) DeletePet(ctx echo.Context, id int64) error {
	panic("implement me")
}

func (api *Api) FindPetById(ctx echo.Context, id int64) error {
	panic("implement me")
}
