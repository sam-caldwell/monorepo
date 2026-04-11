# ------------------------------------------------------------------------------
# Python — multi-stage build
#   builder: Ubuntu 24.04 + Python 3.14 and latest from source + pip3 + venv
#   runtime: distroless with Python (default: 3.14)
# ------------------------------------------------------------------------------

ARG PYTHON_VERSION=3.14

# --- Builder stage ---
FROM opsys/ubuntu:24.04 AS builder

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        build-essential \
        curl \
        ca-certificates \
        libssl-dev \
        zlib1g-dev \
        libbz2-dev \
        libreadline-dev \
        libsqlite3-dev \
        libncursesw5-dev \
        libxml2-dev \
        libxmlsec1-dev \
        libffi-dev \
        liblzma-dev \
        xz-utils && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ARG PYTHON_VERSION

# Install Python 3.14 from source
RUN curl -fsSL "https://www.python.org/ftp/python/3.14.0/Python-3.14.0a3.tar.xz" -o /tmp/python-3.14.tar.xz && \
    mkdir -p /tmp/python-3.14-src && \
    tar -xJf /tmp/python-3.14.tar.xz -C /tmp/python-3.14-src --strip-components=1 && \
    cd /tmp/python-3.14-src && \
    ./configure --prefix=/opt/python/3.14 --enable-optimizations --with-lto && \
    make -j"$(nproc)" && \
    make install && \
    /opt/python/3.14/bin/python3 -m ensurepip --upgrade && \
    /opt/python/3.14/bin/python3 -m pip install --upgrade pip && \
    rm -rf /tmp/python-3.14-src /tmp/python-3.14.tar.xz

# Install latest stable Python from source
RUN LATEST=$(curl -fsSL https://endoflife.date/api/python.json | \
        /opt/python/3.14/bin/python3 -c "import sys,json; print(json.load(sys.stdin)[0]['latest'])") && \
    MAJOR_MINOR=$(echo "$LATEST" | cut -d. -f1-2) && \
    curl -fsSL "https://www.python.org/ftp/python/${LATEST}/Python-${LATEST}.tar.xz" -o /tmp/python-latest.tar.xz && \
    mkdir -p /tmp/python-latest-src && \
    tar -xJf /tmp/python-latest.tar.xz -C /tmp/python-latest-src --strip-components=1 && \
    cd /tmp/python-latest-src && \
    ./configure --prefix="/opt/python/latest" --enable-optimizations --with-lto && \
    make -j"$(nproc)" && \
    make install && \
    /opt/python/latest/bin/python3 -m ensurepip --upgrade && \
    /opt/python/latest/bin/python3 -m pip install --upgrade pip && \
    rm -rf /tmp/python-latest-src /tmp/python-latest.tar.xz

# Create venvs to prove the toolchain works
RUN /opt/python/3.14/bin/python3 -m venv /opt/venv-3.14 && \
    /opt/python/latest/bin/python3 -m venv /opt/venv-latest

WORKDIR /opt/app

# --- Runtime stage ---
FROM opsys/distroless:latest AS runtime

ARG PYTHON_VERSION=3.14

# Copy both Python installations
COPY --from=builder /opt/python /opt/python

# Copy shared libraries Python needs
COPY --from=builder /lib/x86_64-linux-gnu/libz.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libssl.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libcrypto.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libffi.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libbz2.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/liblzma.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libsqlite3.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libreadline.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libncursesw.so.* /lib/x86_64-linux-gnu/
COPY --from=builder /lib/x86_64-linux-gnu/libtinfo.so.* /lib/x86_64-linux-gnu/

ENV PATH="/opt/python/${PYTHON_VERSION}/bin:${PATH}"

WORKDIR /opt/app
USER 65534

ENTRYPOINT ["python3"]
