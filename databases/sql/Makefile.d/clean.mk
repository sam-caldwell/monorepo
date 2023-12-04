clean/sql:
	@color -blue -lf "$@ green"
	@docker exec -it postgresql /bin/bash -c """echo 'drop database tracker;' | /usr/bin/psql""" || true
	@color -green -lf "$@ done"
