#!/usr/bin/env bash
# setup.sh — provision a fresh Debian/Ubuntu VPS as a mehrd validator node.
# Run as root: curl -fsSL https://... | bash
# Or: scp this file to the server and run: bash setup.sh
set -euo pipefail

CHAIN_ID="${CHAIN_ID:-mehr-1}"
MONIKER="${MONIKER:-node-0}"
NODE_HOME="/var/lib/mehrd"
BINARY_URL="${BINARY_URL:-}"   # set to download URL, or leave empty to copy local ./mehrd

# ── 1. System deps ─────────────────────────────────────────────────────────────
apt-get update -qq
apt-get install -y -qq curl wget jq nginx certbot python3-certbot-nginx ufw

# ── 2. Create dedicated user ───────────────────────────────────────────────────
if ! id -u mehrd &>/dev/null; then
  useradd --system --shell /usr/sbin/nologin --home-dir "$NODE_HOME" --create-home mehrd
fi

# ── 3. Install mehrd binary ────────────────────────────────────────────────────
if [ -n "$BINARY_URL" ]; then
  curl -fsSL "$BINARY_URL" -o /usr/local/bin/mehrd
else
  cp ./mehrd /usr/local/bin/mehrd
fi
chmod +x /usr/local/bin/mehrd
chown root:root /usr/local/bin/mehrd

# ── 4. Initialise node ─────────────────────────────────────────────────────────
if [ ! -f "$NODE_HOME/config/genesis.json" ]; then
  sudo -u mehrd mehrd init "$MONIKER" --chain-id "$CHAIN_ID" --home "$NODE_HOME"
fi

# ── 5. Configure node ──────────────────────────────────────────────────────────
# Enable REST API and CORS
CONFIG="$NODE_HOME/config/app.toml"
sed -i 's/^enable = false$/enable = true/'      "$CONFIG"
sed -i 's|^address = "tcp://localhost:1317"|address = "tcp://0.0.0.0:1317"|' "$CONFIG"
sed -i 's/^enabled-unsafe-cors = false/enabled-unsafe-cors = true/' "$CONFIG"

# Enable prometheus (optional)
TMCFG="$NODE_HOME/config/config.toml"
sed -i 's/prometheus = false/prometheus = true/' "$TMCFG"

chown -R mehrd:mehrd "$NODE_HOME"

# ── 6. Systemd service ─────────────────────────────────────────────────────────
cp "$(dirname "$0")/mehrd.service" /etc/systemd/system/mehrd.service
systemctl daemon-reload
systemctl enable mehrd

# ── 7. Firewall ────────────────────────────────────────────────────────────────
ufw allow 22/tcp   # SSH
ufw allow 80/tcp   # HTTP (certbot)
ufw allow 443/tcp  # HTTPS REST
ufw allow 9090/tcp # gRPC
ufw allow 26656/tcp # P2P
ufw --force enable

# ── 8. nginx ───────────────────────────────────────────────────────────────────
DOMAIN="${DOMAIN:-edge-0.mehrfunds.com}"
sed "s/edge-0.mehrfunds.com/$DOMAIN/g" "$(dirname "$0")/nginx-node.conf" \
  > "/etc/nginx/sites-available/$DOMAIN"
ln -sf "/etc/nginx/sites-available/$DOMAIN" "/etc/nginx/sites-enabled/$DOMAIN"
nginx -t && systemctl reload nginx

echo ""
echo "┌──────────────────────────────────────────────────────┐"
echo "│  mehrd installed. Next steps:                        │"
echo "│  1. Copy genesis.json to $NODE_HOME/config/          │"
echo "│  2. Add persistent peers to config.toml              │"
echo "│  3. Run: certbot --nginx -d $DOMAIN                  │"
echo "│  4. Run: systemctl start mehrd                       │"
echo "│  5. Tail logs: journalctl -u mehrd -f                │"
echo "└──────────────────────────────────────────────────────┘"
