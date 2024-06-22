migrate-create:
	@docker exec -it hikarilib-api migrate create -ext sql -dir ./migrations -seq $(name)
	@docker cp hikarilib-api:/go/migrations/. ./migrations
# todo: подтягивать переменные окружения из .env
migrate-up:
	migrate -path ./migrations -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disabled" up
migrate-down:
	migrate -path ./migrations -database "postgres://postgres:qwerty@localhost:5436/postgres?sslmode=disabled" down
remove-image:
	docker rmi hikarilib-backend-hikarilib-api:latest
	docker rmi postgres:13-alpine
up:
	@docker-compose --env-file ./build/.env up
down:
	@docker-compose --env-file ./build/.env down