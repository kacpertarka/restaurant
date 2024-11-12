include .env

current_dir = $(shell pwd)

build:
	docker-compose up --build --remove-orphans

up:
	docker-compose up -d

down:
	docker-compose down

show-logs:
	docker-compose logs

postgres:
	docker-compose exec postgres psql --username=${POSTGRES_USER} --dbname=${POSTGRES_NAME}

migrations:
	docker run -v ./cmd/migrate/migrations:/migrations --network host migrate/migrate create -ext sql -dir /migrations $(filter-out $@,$(MAKECMDGOALS))

migrate-up:
	docker-compose run migrate up

migrate-down:
	docker-compose run migrate down