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

echo:
	echo "lint $(GOLANGCI_LINT)"

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up 1

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down 1

lint: $(GOLANGCI_LINT)
	@$(foreach dir,$(MODULE_DIRS),( \
		cd $(dir) && \
		echo "lint $(GOLANGCI_LINT)" && \
		$(GOLANGCI_LINT) run --config ../.golangci.yml ./...) && \
		) true
