# containers/services/dns-server/Dockerfile
# (c) 2023 Sam Caldwell.  See License.txt
#
# Establish a bind DNS server.
#
FROM opsys/ubuntu:latest AS base

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends bind9 \
                                               bind9utils \
                                               nginx \
                                               openssl

RUN chmod -x /etc/update-motd.d/*
RUN rm -rf /var/www
RUN rm -rf /etc/nginx/*
COPY containers/services/dns-server/src/nginx.conf /etc/nginx/
COPY containers/services/dns-server/src/entrypoint.sh /usr/bin/entrypoint.sh

RUN rm -rf /etc/bind/*
COPY containers/services/dns-server/src/bind/* /etc/bind/

RUN chown -R bind:bind /etc/bind/ && \
    mkdir -p /run/named && \
    chown -R bind:bind /run/named && \
    chmod -R 660 /run/named

WORKDIR root

EXPOSE 80/tcp

EXPOSE 53/udp

ENTRYPOINT [ "/usr/bin/entrypoint.sh" ]

CMD [ "/usr/bin/entrypoint.sh" ]
