FROM opsys/ubuntu:latest AS base

RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y

FROM base AS nodejs

RUN apt-get update -y && \
    apt-get install curl wget xz-utils -y

WORKDIR /tmp
RUN wget https://nodejs.org/dist/v20.10.0/node-v20.10.0-linux-x64.tar.xz && \
    tar -xvJf node-v20.10.0-linux-x64.tar.xz && \
    cd node-v20.10.0-linux-x64

FROM nodejs AS runtime
WORKDIR /tmp
COPY --from=nodejs /tmp/node-v20.10.0-linux-x64/bin /usr/local/bin/
COPY --from=nodejs /tmp/node-v20.10.0-linux-x64/include /usr/local/include/
COPY --from=nodejs /tmp/node-v20.10.0-linux-x64/lib /usr/local/lib/
COPY --from=nodejs /tmp/node-v20.10.0-linux-x64/share /usr/local/share/
