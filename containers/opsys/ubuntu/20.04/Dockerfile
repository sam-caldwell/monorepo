FROM ubuntu:20.04

RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends ca-certificates apt-transport-https

WORKDIR /opt/

# To Do: Create user account (service user)
# To Do: Install any security tooling, etc.
