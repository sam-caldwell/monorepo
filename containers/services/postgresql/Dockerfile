# containers/services/postgresql/Dockerfile
# (c) 2023 Sam Caldwell.  See License.txt
#
# We need a postgresql server
#
FROM opsys/ubuntu:latest AS base

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends postgresql

RUN echo "listen_addresses = '0.0.0.0'" >> /etc/postgresql/15/main/postgresql.conf
RUN echo "log_timezone = 'America/Chicago'" >> /etc/postgresql/15/main/postgresql.conf
RUN echo "autovacuum = on" >> /etc/postgresql/15/main/postgresql.conf
#RUN echo "local   all             postgres                                peer" > /etc/postgresql/15/main/pg_hba.conf
RUN echo "host    all             all      0.0.0.0/0  scram-sha-256" >> /etc/postgresql/15/main/pg_hba.conf

COPY containers/services/postgresql/src/create_admin_role.sql /opt/
USER postgres
RUN /etc/init.d/postgresql start &&  \
    echo 'waiting on db to start' && \
    /etc/init.d/postgresql status && \
    sleep 5 && \
    echo 'database should be started' && \
    /usr/bin/psql < /opt/create_admin_role.sql
USER root
RUN rm /opt/create_admin_role.sql
RUN /etc/init.d/postgresql stop

USER postgres
WORKDIR /var/lib/postgresql
VOLUME ["/var/lib/postgresql/15/main"]
EXPOSE 5432

ENTRYPOINT ["/usr/lib/postgresql/15/bin/postgres", "-c",  "config_file=/etc/postgresql/15/main/postgresql.conf"]
