# ------------------------------------------------------------------------------
# Java 21 — multi-stage build
#   builder: Ubuntu 24.04 + OpenJDK 21 (full JDK for compilation)
#   runtime: distroless + JRE 21 only
# ------------------------------------------------------------------------------

# --- Builder stage ---
FROM opsys/ubuntu:24.04 AS builder

RUN apt-get update -y && \
    apt-get install -y --no-install-recommends \
        openjdk-21-jdk-headless \
        binutils && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

ENV JAVA_HOME=/usr/lib/jvm/java-21-openjdk-amd64
ENV PATH="${JAVA_HOME}/bin:${PATH}"

# Build a minimal JRE with jlink for the runtime stage
RUN jlink \
    --add-modules ALL-MODULE-PATH \
    --strip-debug \
    --no-man-pages \
    --no-header-files \
    --compress=zip-6 \
    --output /opt/jre-minimal

WORKDIR /opt/app

# --- Runtime stage ---
FROM opsys/distroless:latest AS runtime

COPY --from=builder /opt/jre-minimal /opt/jre
COPY --from=builder /lib/x86_64-linux-gnu/libz.so.* /lib/x86_64-linux-gnu/

ENV JAVA_HOME=/opt/jre
ENV PATH="${JAVA_HOME}/bin:${PATH}"

WORKDIR /opt/app
USER 65534

ENTRYPOINT ["java"]
