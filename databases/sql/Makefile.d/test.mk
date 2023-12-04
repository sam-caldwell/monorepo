# sql/Makefile.d/test.mk
# (c) 2023 Sam Caldwell.  See License.txt
#
test/sql:
	@color -blue -lf -reset "start $@"
	@db_host=192.168.3.190 \
	db_port=5432 \
	db_name=tracker \
	db_user=admin \
	db_pass=admin  \
	use_tls=false go test -v -failfast ./databases/sql/... || {\
	  color -red -lf -reset"failed $@";\
	  exit 1;\
  	};\
	color -green -lf -reset "completed $@"

