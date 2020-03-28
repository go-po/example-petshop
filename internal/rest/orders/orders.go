package orders

import (
	"context"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-po/example-petstore/generated/server/models"
	"github.com/go-po/example-petstore/generated/server/restapi/operations"
	"github.com/go-po/example-petstore/generated/server/restapi/operations/store"
	"github.com/go-po/example-petstore/internal/domain/orders"
	"time"
)

type App interface {
	PlaceOrder(ctx context.Context, cmd orders.PlaceOrderCmd) error
}

func Register(api *operations.ExamplePetstoreAPI, app App) {
	api.StorePlaceOrderHandler = placeOrder(app)
	api.StoreGetOrderByIDHandler = getOrdersById(app)
}

func placeOrder(app App) store.PlaceOrderHandlerFunc {
	return func(params store.PlaceOrderParams) middleware.Responder {
		err := app.PlaceOrder(params.HTTPRequest.Context(), orders.PlaceOrderCmd{
			ID:       params.Body.ID,
			PetID:    params.Body.PetID,
			Quantity: params.Body.Quantity,
			ShipDate: time.Time(params.Body.ShipDate),
		})
		if err != nil {
			return store.NewPlaceOrderBadRequest()
		}
		return store.NewPlaceOrderOK().WithPayload(&models.Order{
			Complete: false,
			PetID:    params.Body.PetID,
			Quantity: params.Body.Quantity,
			ID:       params.Body.ID,
			ShipDate: params.Body.ShipDate,
			Status:   "placed",
		})
	}
}

func getOrdersById(app App) store.GetOrderByIDHandlerFunc {
	return func(params store.GetOrderByIDParams) middleware.Responder {
		return store.NewGetOrderByIDBadRequest()
	}
}
