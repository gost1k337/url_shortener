RUNNER=migrate

ifeq ($(p),host)
 	RUNNER=sql-migrate
endif

SOURCE="FILE://MIGRATIONS"
MIGRATE=$(RUNNER)
DB=${POSTGRES_DSN}

migrate-status:
	$(MIGRATE) version

migrate-up:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" up 1

migrate-down:
	$(MIGRATE) -source ${SOURCE} -database "${DB}" down 1