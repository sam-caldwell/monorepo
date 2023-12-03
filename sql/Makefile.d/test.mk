# sql/Makefile.d/test.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/sql:
	@echo ${GREEN}"start $@"${RESET}
	@db_host=192.168.3.190 \
	db_port=5432 \
	db_name=tracker \
	db_user=admin \
	db_pass=admin  \
	use_tls=false go test -v -failfast ./sql/databases/... || {\
	  echo ${RED}"failed $@"${RESET};\
	  exit 1;\
  	};\
	echo ${GREEN}"completed $@"${RESET}

