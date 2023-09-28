include .env
export

SWAGGER=swag
SWAGGER_PATH="./internal/app/app.go"

swagger init:
	$(SWAGGER) init -g ${SWAGGER_PATH}

build:
	docker-compose -p ags build

run:
	docker-compose -p ags up -d

stop:
	docker-compose -p ags stop

run-local:
	go run cmd/app/main.go
