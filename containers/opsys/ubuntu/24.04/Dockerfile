FROM ubuntu:24.04

ENV DEBIAN_FRONTEND=noninteractive

# --- Base packages ---
RUN apt-get update -y --fix-missing && \
    apt-get upgrade -y && \
    apt-get install -y --no-install-recommends \
        ca-certificates \
        apt-transport-https && \
    apt-get autoremove -y && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

# --- CIS Benchmark hardening (Ubuntu 24.04 Level 1 - Server) ---

# 1.1 - Ensure unnecessary filesystems are disabled
RUN for fs in cramfs freevxfs jffs2 hfs hfsplus squashfs udf; do \
        echo "install ${fs} /bin/true" >> /etc/modprobe.d/cis-disable-filesystems.conf; \
    done

# 1.5 - Restrict core dumps
RUN echo '* hard core 0' >> /etc/security/limits.conf && \
    echo 'fs.suid_dumpable = 0' >> /etc/sysctl.d/99-cis.conf

# 3.1 - Network hardening via sysctl
RUN { \
    echo 'net.ipv4.ip_forward = 0'; \
    echo 'net.ipv4.conf.all.send_redirects = 0'; \
    echo 'net.ipv4.conf.default.send_redirects = 0'; \
    echo 'net.ipv4.conf.all.accept_source_route = 0'; \
    echo 'net.ipv4.conf.default.accept_source_route = 0'; \
    echo 'net.ipv4.conf.all.accept_redirects = 0'; \
    echo 'net.ipv4.conf.default.accept_redirects = 0'; \
    echo 'net.ipv4.conf.all.secure_redirects = 0'; \
    echo 'net.ipv4.conf.default.secure_redirects = 0'; \
    echo 'net.ipv4.conf.all.log_martians = 1'; \
    echo 'net.ipv4.conf.default.log_martians = 1'; \
    echo 'net.ipv4.icmp_echo_ignore_broadcasts = 1'; \
    echo 'net.ipv4.icmp_ignore_bogus_error_responses = 1'; \
    echo 'net.ipv4.conf.all.rp_filter = 1'; \
    echo 'net.ipv4.conf.default.rp_filter = 1'; \
    echo 'net.ipv4.tcp_syncookies = 1'; \
    echo 'net.ipv6.conf.all.accept_ra = 0'; \
    echo 'net.ipv6.conf.default.accept_ra = 0'; \
    echo 'net.ipv6.conf.all.accept_redirects = 0'; \
    echo 'net.ipv6.conf.default.accept_redirects = 0'; \
    } >> /etc/sysctl.d/99-cis.conf

# 5.2 - SSH hardening (sshd defaults for images that install openssh later)
RUN mkdir -p /etc/ssh/sshd_config.d && \
    { \
    echo 'Protocol 2'; \
    echo 'LogLevel VERBOSE'; \
    echo 'MaxAuthTries 4'; \
    echo 'PermitRootLogin no'; \
    echo 'PermitEmptyPasswords no'; \
    echo 'IgnoreRhosts yes'; \
    echo 'HostbasedAuthentication no'; \
    echo 'X11Forwarding no'; \
    echo 'ClientAliveInterval 300'; \
    echo 'ClientAliveCountMax 3'; \
    echo 'LoginGraceTime 60'; \
    echo 'Banner /etc/issue.net'; \
    } > /etc/ssh/sshd_config.d/99-cis.conf

# 5.4 - Password and login policies
RUN { \
    echo 'PASS_MAX_DAYS 365'; \
    echo 'PASS_MIN_DAYS 7'; \
    echo 'PASS_WARN_AGE 7'; \
    } >> /etc/login.defs && \
    echo 'umask 027' >> /etc/profile.d/cis-umask.sh

# 1.4 / 6.1 - File permission hardening
RUN chmod 644 /etc/passwd /etc/group && \
    chmod 640 /etc/shadow 2>/dev/null || true && \
    chmod 600 /etc/ssh/sshd_config.d/99-cis.conf

# 6.2 - Remove unnecessary SUID/SGID binaries
RUN find / -xdev -type f \( -perm -4000 -o -perm -2000 \) -exec chmod a-s {} + 2>/dev/null || true

# Warning banner
RUN echo 'Authorized uses only. All activity may be monitored and reported.' > /etc/issue.net

# --- Non-root service user ---
RUN groupadd -r app && useradd -r -g app -d /opt/app -s /usr/sbin/nologin app

WORKDIR /opt/app
