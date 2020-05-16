package main

import (
	"log"

	"github.com/go-po/example-petstore/generated/petstore"
	"github.com/go-po/example-petstore/internal/domain/pets"
	"github.com/go-po/example-petstore/internal/rest"
	"github.com/go-po/po"
	"github.com/labstack/echo/v4"
)

func main() {

	es, err := po.NewFromOptions(
		po.WithStoreInMemory(),
		po.WithProtocolChannels(),
	)
	if err != nil {
		log.Fatalf("event source: %s", err)
	}

	petsApp, err := pets.New(es)
	if err != nil {
		log.Fatalf("pets app: %s", err)
	}

	e := echo.New()
	api := rest.New(petsApp)
	petstore.RegisterHandlers(e, api)

	if err := e.Start("localhost:8080"); err != nil {
		log.Fatalln(err)
	}

}
