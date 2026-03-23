#!/usr/bin/env bash
set -euo pipefail

mkdir -p /usr/local/etc/uuhappyserver
UUID=$(/usr/local/bin/uuhappyserver uuid)
KEYS=$(/usr/local/bin/uuhappyserver x25519)
PRIV=$(echo "$KEYS" | awk -F': ' '/PrivateKey/ {print $2}')
PUB=$(echo "$KEYS" | awk -F': ' '/PublicKey/ {print $2}')

cat > /usr/local/etc/uuhappyserver/config.json <<EOF
{
  "log": {
    "loglevel": "warning"
  },
  "inbounds": [
    {
      "listen": "0.0.0.0",
      "port": 443,
      "protocol": "vless",
      "settings": {
        "clients": [
          {
            "id": "$UUID",
            "flow": "xtls-rprx-vision"
          }
        ],
        "decryption": "none"
      },
      "streamSettings": {
        "network": "tcp",
        "security": "reality",
        "realitySettings": {
          "show": false,
          "dest": "www.cloudflare.com:443",
          "xver": 0,
          "serverNames": [
            "www.cloudflare.com"
          ],
          "privateKey": "$PRIV",
          "shortIds": [
            ""
          ]
        }
      },
      "sniffing": {
        "enabled": true,
        "destOverride": [
          "http",
          "tls",
          "quic"
        ]
      }
    }
  ],
  "outbounds": [
    {
      "protocol": "freedom"
    },
    {
      "protocol": "blackhole",
      "tag": "blocked"
    }
  ]
}
EOF

cat > /etc/systemd/system/uuhappyserver.service <<'EOF'
[Unit]
Description=uuhappyserver Service
After=network.target nss-lookup.target
Wants=network.target

[Service]
Type=simple
User=root
CapabilityBoundingSet=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
AmbientCapabilities=CAP_NET_ADMIN CAP_NET_BIND_SERVICE
NoNewPrivileges=true
ExecStart=/usr/local/bin/uuhappyserver run -c /usr/local/etc/uuhappyserver/config.json
Restart=on-failure
RestartSec=5s
LimitNOFILE=1048576

[Install]
WantedBy=multi-user.target
EOF

/usr/local/bin/uuhappyserver run -test -c /usr/local/etc/uuhappyserver/config.json

systemctl stop nginx || true
systemctl disable nginx || true
systemctl daemon-reload
systemctl enable --now uuhappyserver

sleep 1

echo "UUID=$UUID"
echo "PUBLIC_KEY=$PUB"
echo "REALITY_SERVERNAME=www.cloudflare.com"
echo "PORT=443"
echo "SERVICE_STATUS=$(systemctl is-active uuhappyserver)"
ss -lntp | grep ':443' || true
