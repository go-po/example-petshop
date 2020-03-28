package rest

import (
	"github.com/go-openapi/loads"
	"github.com/go-po/example-petstore/generated/server/restapi"
	"github.com/go-po/example-petstore/generated/server/restapi/operations"
	"github.com/go-po/example-petstore/internal/rest/orders"
	"github.com/go-po/example-petstore/internal/rest/pets"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

func New(ordersApp orders.App, petsApp pets.App) (*restapi.Server, error) {
	swaggerSpec, err := loads.Embedded(restapi.SwaggerJSON, restapi.FlatSwaggerJSON)
	if err != nil {
		log.Fatalln(err)
	}

	api := operations.NewExamplePetstoreAPI(swaggerSpec)

	orders.Register(api, ordersApp)
	pets.Register(api, petsApp)

	server := restapi.NewServer(api)

	parser := flags.NewParser(server, flags.Default)
	parser.ShortDescription = "Swagger Petstore"
	parser.LongDescription = "This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters."
	server.ConfigureFlags()
	for _, optsGroup := range api.CommandLineOptionsGroups {
		_, err := parser.AddGroup(optsGroup.ShortDescription, optsGroup.LongDescription, optsGroup.Options)
		if err != nil {
			return nil, err
		}
	}

	if _, err := parser.Parse(); err != nil {
		code := 1
		if fe, ok := err.(*flags.Error); ok {
			if fe.Type == flags.ErrHelp {
				code = 0
			}
		}
		os.Exit(code)
	}

	server.ConfigureAPI()

	return server, nil
}
