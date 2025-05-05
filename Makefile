include .env

up:
	docker compose up -d

stop:
	docker stop $(docker ps -q)

build:
	docker compose build --no-cache

down:
	docker compose down -v

clean:
	docker system prune -a --volumes -f && docker builder prune -a -f

logs:
	docker compose logs -f

start: build up

end: down clean

# Execute this on Git Bash or other Unix Shell
list:
	export $(shell cat .env | sed 's/#.*//;/^$$/d' | xargs) && docker exec -it pg psql -U $(DB_USER) -d $(DB_NAME) -c "\dt"

gin:
	docker exec -it gin bash

pg:
	docker exec -it pg bash

swag:
	swag init -g cmd/api/main.go -o ./docs