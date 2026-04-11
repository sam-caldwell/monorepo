# ------------------------------------------------------------------------------
# Node.js — multi-stage build
#   builder: Ubuntu 24.04 + nvm with Node 20, 21, and latest + npm + yarn
#   runtime: distroless with Node (default: latest)
# ------------------------------------------------------------------------------

ARG NODE_VERSION=latest

# --- Builder stage ---
FROM opsys/ubuntu:24.04 AS builder

ENV NVM_DIR=/opt/nvm
ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        curl \
        ca-certificates \
        xz-utils && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

RUN mkdir -p "${NVM_DIR}" && \
    curl -fsSL https://raw.githubusercontent.com/nvm-sh/nvm/v0.40.1/install.sh | bash

SHELL ["/bin/bash", "-c"]

# Install Node versions: 20 (LTS), 21, and latest
RUN source "${NVM_DIR}/nvm.sh" && \
    nvm install 20 && \
    nvm install 21 && \
    nvm install node && \
    nvm alias default node

# Install yarn globally for each version
RUN source "${NVM_DIR}/nvm.sh" && \
    for v in 20 21 node; do \
        nvm use "$v" && \
        npm install -g yarn; \
    done

WORKDIR /opt/app

# --- Runtime stage ---
FROM opsys/distroless:latest AS runtime

ARG NODE_VERSION=latest

COPY --from=builder /opt/nvm /opt/nvm

# Copy the shared libraries Node.js needs
COPY --from=builder /lib/x86_64-linux-gnu/libstdc++.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libgcc_s.so.* /lib/x86_64-linux-gnu/

# Resolve the real node version directory and symlink it
# We copy a small shell helper that the builder prepares
SHELL ["/bin/bash", "-c"]

ENV NVM_DIR=/opt/nvm
ENV PATH="/opt/nvm/versions/node/default/bin:${PATH}"

WORKDIR /opt/app
USER 65534

ENTRYPOINT ["node"]
