clean/sql:
	@color -blue -lf -reset "$@ start"
	@docker exec -it postgresql /bin/bash -c """echo 'drop database tracker;' | /usr/bin/psql""" &>/dev/null || true
	@color -green -lf -reset "$@ done"
