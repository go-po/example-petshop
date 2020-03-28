package pets

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-po/example-petstore/generated/server/models"
	"github.com/go-po/example-petstore/generated/server/restapi/operations"
	"github.com/go-po/example-petstore/generated/server/restapi/operations/pet"
)

type App interface {
}

func Register(api *operations.ExamplePetstoreAPI, app App) {
	api.PetAddPetHandler = addPet(app)
	api.PetDeletePetHandler = deletePet(app)
	api.PetFindPetsByStatusHandler = findPetByStatus(app)
	api.PetFindPetsByTagsHandler = findPetByTags(app)
	api.PetGetPetByIDHandler = getPetById(app)
	api.PetUpdatePetHandler = updatePet(app)
	api.PetUpdatePetWithFormHandler = updatePetWithForm(app)
	api.PetUploadFileHandler = uploadFile(app)
}

func uploadFile(app App) pet.UploadFileHandlerFunc {
	return func(params pet.UploadFileParams, principle interface{}) middleware.Responder {
		return pet.NewUploadFileOK().WithPayload(&models.APIResponse{
			Code:    0,
			Message: "",
			Type:    "",
		})
	}
}

func updatePetWithForm(app App) pet.UpdatePetWithFormHandlerFunc {
	return func(params pet.UpdatePetWithFormParams, principle interface{}) middleware.Responder {
		return pet.NewUpdatePetBadRequest()
	}
}

func updatePet(app App) pet.UpdatePetHandlerFunc {
	return func(params pet.UpdatePetParams, principle interface{}) middleware.Responder {
		return pet.NewUpdatePetBadRequest()
	}
}

func getPetById(app App) pet.GetPetByIDHandlerFunc {
	return func(params pet.GetPetByIDParams, principle interface{}) middleware.Responder {
		return pet.NewGetPetByIDBadRequest()
	}
}

func findPetByTags(app App) pet.FindPetsByTagsHandlerFunc {
	return func(params pet.FindPetsByTagsParams, principle interface{}) middleware.Responder {
		return pet.NewFindPetsByTagsBadRequest()
	}
}

func findPetByStatus(app App) pet.FindPetsByStatusHandlerFunc {
	return func(params pet.FindPetsByStatusParams, principle interface{}) middleware.Responder {
		return pet.NewFindPetsByStatusBadRequest()
	}
}

func deletePet(app App) pet.DeletePetHandlerFunc {
	return func(params pet.DeletePetParams, principle interface{}) middleware.Responder {
		return pet.NewDeletePetBadRequest()
	}
}

func addPet(app App) pet.AddPetHandlerFunc {
	return func(params pet.AddPetParams, principle interface{}) middleware.Responder {
		return pet.NewAddPetMethodNotAllowed()
	}
}
