build:
	docker-compose up --build --remove-orphans

up:
	docker-compose up -d

down:
	docker-compose down

show-logs:
	docker-compose logs

postgres:
	docker-compose exec postgres psql --username=postgres --dbname=postgres