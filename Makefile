SWAGGER=docker run --rm -it -e GOPATH=${HOME}/go:/go -v ${HOME}:${HOME} -w $(shell pwd) quay.io/goswagger/swagger

stuff:
	$(SWAGGER) generate server --help

gen-server:
	$(SWAGGER) generate server 	-f ./api/petstore.yaml \
							  	-t generated/server \
  								--exclude-spec \
  								--exclude-main \
							  	-A "Example Petstore"

db-reset:
	echo "drop schema public cascade" 			| docker exec -i petstore_pg psql -U petstore po
	echo "create schema if not exists public" 	| docker exec -i petstore_pg psql -U petstore petstore

mq-reset:
	./scripts/reset-rabbit.sh

reset: db-reset mq-reset

test:
	go test ./... -count 1 -race

test-short:
	go test ./... -count 1 -race -test.short

cover:
	go test ./... -count 1 -race -cover -test.short

cover-all:
	go test ./... -count 1 -race -cover

gen: clean
	go generate ./...

up:
	docker-compose up -d

down:
	docker-compose down

plantuml:
	./scripts/plantuml.sh

open-mq:
	open http://localhost:15669/ -a Safari
