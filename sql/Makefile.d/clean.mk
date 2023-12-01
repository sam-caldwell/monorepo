clean/sql:
	@docker exec -it postgresql /bin/bash -c """echo 'drop database tracker;' | /usr/bin/psql""" || true

