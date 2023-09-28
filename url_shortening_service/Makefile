include .env
export

.PHONY: protos

protos:
	protoc -I api/protos/ --go_out=./api/protos --go-grpc_out=./api/protos api/protos/url_shorts.proto

build:
	docker-compose -p uss build

run:
	docker-compose -p uss up -d

stop:
	docker-compose -p uss stop

run-local:
	go run cmd/app/main.go
