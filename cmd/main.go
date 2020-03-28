package main

import (
	"github.com/go-po/example-petstore/internal/domain/orders"
	"github.com/go-po/example-petstore/internal/domain/pets"
	"github.com/go-po/example-petstore/internal/domain/users"
	"github.com/go-po/example-petstore/internal/rest"
	"github.com/go-po/po"
	"log"
)

func main() {

	es, err := po.NewFromOptions(
		po.WithStoreInMemory(),
		po.WithProtocolChannels(),
	)
	if err != nil {
		log.Fatalf("event source: %s", err)
	}

	ordersApp, err := orders.New(es)
	if err != nil {
		log.Fatalf("orders app: %s", err)
	}

	petsApp, err := pets.New(es)
	if err != nil {
		log.Fatalf("pets app: %s", err)
	}

	usersApp, err := users.New(es)
	if err != nil {
		log.Fatalf("users app: %s", err)
	}

	server, err := rest.New(ordersApp, petsApp, usersApp)
	if err != nil {
		log.Fatalf("rest serverr: %s", err)
	}
	defer server.Shutdown()

	if err := server.Serve(); err != nil {
		log.Fatalln(err)
	}

}
