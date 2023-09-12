SHELL:=/bin/bash

run:
	go run main.go

docker-rebuild:
	docker-compose -f ./deploy/docker-compose.yaml build user

docker-rebuild-tr:
	docker-compose -f ./deploy/docker-compose.yaml build transactions

docker-run:
	docker-compose -f ./deploy/docker-compose.yaml up

db-run:
	docker-compose -f ./deploy/docker-compose.yaml up --remove-orphans user-db-postgres

nats-run:
	docker-compose -f ./deploy/docker-compose.yaml up --remove-orphans nats

migration-build:
	go build -v -o ./bin/migrations ./migrations/

migration-create-sql:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations create $(name) sql

migration-create-go:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations create $(name) go

migration-status:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations status

migration-up:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations up

migration-up-by-one:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations up-by-one

migration-down-to:
	go build -v -o ./bin/migrations ./migrations/ && \
    ./bin/migrations down-to $(version)
