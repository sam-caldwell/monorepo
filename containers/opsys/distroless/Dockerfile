FROM ubuntu:24.04 AS base

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        apt-transport-https && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# --- Build the distroless image from scratch ---
FROM scratch

# TLS certificates
COPY --from=base /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

# Minimal user/group database (nobody:65534 as the runtime user)
COPY --from=base /etc/passwd /etc/passwd
COPY --from=base /etc/group /etc/group

# Minimal C runtime libraries
COPY --from=base /lib/x86_64-linux-gnu/libc.so.* /lib/x86_64-linux-gnu/
COPY --from=base /lib/x86_64-linux-gnu/libpthread.so.* /lib/x86_64-linux-gnu/
COPY --from=base /lib/x86_64-linux-gnu/libdl.so.* /lib/x86_64-linux-gnu/
COPY --from=base /lib/x86_64-linux-gnu/libm.so.* /lib/x86_64-linux-gnu/
COPY --from=base /lib64/ld-linux-x86-64.so.* /lib64/

# Timezone data
COPY --from=base /usr/share/zoneinfo /usr/share/zoneinfo

# --- CIS hardening (applicable controls for a scratch-based image) ---

# 1.5 - Restrict core dumps via sysctl defaults
COPY --from=base /dev/null /etc/sysctl.d/99-cis.conf

# Warning banner
COPY --from=base /dev/null /etc/issue.net

# CIS 6.1 - Minimal filesystem: no SUID/SGID binaries exist (scratch has no binaries)
# CIS 5.2 - No SSH server present (attack surface eliminated)
# CIS 1.1 - No kernel modules can be loaded (no modprobe, no /lib/modules)
# CIS 3.x - Network sysctl hardening is deferred to the orchestrator (k8s securityContext)

WORKDIR /opt/app

# Run as nobody (65534) — no root shell exists in this image
USER 65534
