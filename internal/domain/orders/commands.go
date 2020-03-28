package orders

import "time"

type PlaceOrderCmd struct {
	ID       int64
	PetID    int64
	Quantity int32
	ShipDate time.Time
}
