include .env
export

.PHONY: protos

protos:
	protoc -I api/protos/ --go_out=./api/protos --go-grpc_out=./api/protos api/protos/url_shorts.proto

run dev:
	go run cmd/app/main.go