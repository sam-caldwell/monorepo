# sql/Makefile.d/build.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
build/sql: build/go/db/dbMigrations
	@echo ${GREEN}"$@ starting"${RESET}

	@echo ${GREEN}"Create the database schema in the local dev database"${RESET}
	@docker kill postgresql &>/dev/null || true
	@docker rm postgresql &>/dev/null || true
	@docker run -d --rm --name postgresql -p 0.0.0.0:5432:5432 -v ./sql/databases:/opt/ services/postgresql
	@docker ps
	@echo ${GREEN}"$@ has started the local dev database"${RESET}

	go run go/db/dbMigrations/main.go
	@echo ${GREEN}"$@ completed"${RESET}
