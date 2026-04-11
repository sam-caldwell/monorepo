# ------------------------------------------------------------------------------
# PostgreSQL 16 — multi-stage build
#   ubuntu:  Ubuntu 24.04 + PostgreSQL 16 (full server for dev/CI)
#   runtime: distroless + PostgreSQL 16 binaries and libs only
# ------------------------------------------------------------------------------

# --- Builder stage: compile from the PostgreSQL APT repo ---
FROM opsys/ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive
ARG PG_MAJOR=16

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        curl \
        ca-certificates \
        gnupg \
        lsb-release && \
    echo "deb http://apt.postgresql.org/pub/repos/apt $(lsb_release -cs)-pgdg main" \
        > /etc/apt/sources.list.d/pgdg.list && \
    curl -fsSL https://www.postgresql.org/media/keys/ACCC4CF8.asc | gpg --dearmor \
        -o /etc/apt/trusted.gpg.d/postgresql.gpg && \
    apt-get update -y && \
    apt-get install -y --no-install-recommends \
        "postgresql-${PG_MAJOR}" && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# --- Ubuntu target: full server image ---
FROM builder AS ubuntu

ARG PG_MAJOR=16

RUN echo "listen_addresses = '0.0.0.0'" >> "/etc/postgresql/${PG_MAJOR}/main/postgresql.conf" && \
    echo "log_timezone = 'UTC'" >> "/etc/postgresql/${PG_MAJOR}/main/postgresql.conf" && \
    echo "autovacuum = on" >> "/etc/postgresql/${PG_MAJOR}/main/postgresql.conf" && \
    echo "host all all 0.0.0.0/0 scram-sha-256" >> "/etc/postgresql/${PG_MAJOR}/main/pg_hba.conf"

COPY containers/services/postgresql/src/create_admin_role.sql /tmp/
USER postgres
RUN pg_ctlcluster ${PG_MAJOR} main start && \
    sleep 3 && \
    psql < /tmp/create_admin_role.sql && \
    pg_ctlcluster ${PG_MAJOR} main stop
USER root
RUN rm /tmp/create_admin_role.sql

USER postgres
WORKDIR /var/lib/postgresql
VOLUME ["/var/lib/postgresql/${PG_MAJOR}/main"]
EXPOSE 5432

ENTRYPOINT ["/usr/lib/postgresql/16/bin/postgres", "-c", "config_file=/etc/postgresql/16/main/postgresql.conf"]

# --- Distroless target: minimal runtime ---
FROM opsys/distroless:latest AS runtime

ARG PG_MAJOR=16

COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/bin/postgres /usr/lib/postgresql/${PG_MAJOR}/bin/postgres
COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/bin/pg_ctl /usr/lib/postgresql/${PG_MAJOR}/bin/pg_ctl
COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/bin/initdb /usr/lib/postgresql/${PG_MAJOR}/bin/initdb
COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/lib/ /usr/lib/postgresql/${PG_MAJOR}/lib/
COPY --from=builder /usr/lib/postgresql/${PG_MAJOR}/share/ /usr/lib/postgresql/${PG_MAJOR}/share/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libpq.so* /usr/lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libssl.so* /usr/lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libcrypto.so* /usr/lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libreadline.so* /usr/lib/x86_64-linux-gnu/
COPY --from=builder /usr/lib/x86_64-linux-gnu/libldap.so* /usr/lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libz.so* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/liblzma.so* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libtinfo.so* /lib/x86_64-linux-gnu/
COPY --from=builder /etc/postgresql/${PG_MAJOR}/ /etc/postgresql/${PG_MAJOR}/

ENV PATH="/usr/lib/postgresql/${PG_MAJOR}/bin:${PATH}"

WORKDIR /var/lib/postgresql
VOLUME ["/var/lib/postgresql/data"]
EXPOSE 5432

ENTRYPOINT ["postgres"]
