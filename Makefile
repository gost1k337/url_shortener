include .env
export

RUNNER=migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

SOURCE="FILE://MIGRATIONS"
MIGRATE=$(RUNNER)
DB=${POSTGRES_DSN}

GOLANGCI_LINT = /opt/homebrew/bin/golangci-lint
MODULE_DIRS = ./user_service ./api_gateway_service ./url_shortening_service

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down

lint: $(GOLANGCI_LINT)
	@$(foreach dir,$(MODULE_DIRS),( \
		cd $(dir) && \
		$(GOLANGCI_LINT) run --config ../.golangci.yml ./...) && \
		) true
