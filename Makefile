
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
