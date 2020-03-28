package pets

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-po/example-petstore/generated/server/restapi/operations"
	"github.com/go-po/example-petstore/generated/server/restapi/operations/pet"
)

type App interface {
}

func Register(api *operations.ExamplePetstoreAPI, app App) {
	api.PetAddPetHandler = addPet(app)
}

func addPet(app App) pet.AddPetHandlerFunc {
	return func(params pet.AddPetParams, principle interface{}) middleware.Responder {
		return pet.NewAddPetMethodNotAllowed()
	}
}
