include .env
export

.PHONY: protos

protos:
	protoc -I api/protos/ --go_out=./api/protos --go-grpc_out=./api/protos api/protos/user.proto

build:
	docker-compose -p us build

run:
	docker-compose -p us up -d

stop:
	docker-compose -p us stop

run-local:
	go run cmd/app/main.go
