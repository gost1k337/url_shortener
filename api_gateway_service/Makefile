include .env
export

SWAGGER=swag
SWAGGER_PATH="./internal/app/app.go"

swagger init:
	$(SWAGGER) init -g ${SWAGGER_PATH}

run dev:
	go run cmd/app/main.go