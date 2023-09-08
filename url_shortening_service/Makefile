include .env
export

RUNNER=migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

.PHONY: protos
SOURCE="FILE://MIGRATIONS"
MIGRATE=$(RUNNER)
DB=${POSTGRES_DSN}

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up 1

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down 1

protos:
	protoc -I api/protos/ --go_out=./api/protos --go-grpc_out=./api/protos api/protos/url_shorts.proto

run dev:
	go run cmd/app/main.go