# sql/Makefile.d/build.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
build/sql: build/go/db/dbMigrations
	@color -blue -reset -lf "$@ starting"
	@color -blue -reset -lf "Create the database schema in the local dev database"
	@docker kill postgresql &>/dev/null || true
	@docker rm postgresql &>/dev/null || true
	@docker run -d --rm --name postgresql -p 0.0.0.0:5432:5432 -v ./databases/sql:/opt/ services/postgresql || {\
			color -red -lf -reset ">>>failed $@";\
			exit 1;\
	};\
	color -green -lf -reset "<<complete $@";\
	docker ps;\
	color -blue -reset -lf $@ has started the local dev database;\
	go run go/db/dbMigrations/main.go --db_host 192.168.3.190 \
									  --db_port 5432 \
									  --db_user admin \
									  --db_pass admin || {\
			color -red -lf -reset ">>>failed $@";\
			exit 1;\
	}; \
	color -green -reset -lf "$@ completed"
