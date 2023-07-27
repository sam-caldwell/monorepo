#!/bin/sh -eux

motd_content="created $(date)"

if [ -d /etc/update-motd.d ]; then
    MOTD_CONFIG='/etc/update-motd.d/99-created'

    cat >> "$MOTD_CONFIG" <<BENTO
#!/bin/sh

cat <<'EOF'
$motd_content
EOF
BENTO

    chmod 0755 "$MOTD_CONFIG"
else
    echo "${motd_content}" >> /etc/motd
fi
