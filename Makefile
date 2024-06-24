ifneq (,$(wildcard ./build/.env))
    include ./build/.env
    export $(shell sed 's/=.*//' ./build/.env)
endif

migrate-create:
	@docker exec -it hikarilib-api migrate create -ext sql -dir ./migrations -seq $(name)
migrate-up:
	migrate -path ./migrations -database "$(DB_SERVICE)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" up
migrate-down:
	migrate -path ./migrations -database "$(DB_SERVICE)://$(DB_USERNAME):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" down
remove-image:
	docker rmi hikarilib-backend-hikarilib-api:latest
	docker rmi postgres:13-alpine
up:
	@docker-compose --env-file ./build/.env up
down:
	@docker-compose --env-file ./build/.env down