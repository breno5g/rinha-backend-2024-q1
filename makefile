.PHONY: default run build clean

APP_NAME=rinha

default: run

run:
	@go run ./cmd/server/main.go

build:
	@go build -o ${APP_NAME} ./cmd/server/main.go

compose-up-local:
	@docker-compose -f docker-compose-local.yml up

compose-down-local:
	@docker-compose -f docker-compose-local.yml down

compose-up-prod:
	@docker-compose -f docker-compose.yml up

compose-down-prod:
	@docker-compose -f docker-compose.yml down

clean:
	@rm -f ${APP_NAME}
	@docker-compose -f docker-compose-local.yml down
	@docker-compose -f docker-compose.yml down