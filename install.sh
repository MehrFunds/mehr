#!/bin/sh
# Mehr daemon installer and admin panel
# Usage: curl -fsSL https://raw.githubusercontent.com/MehrFunds/mehr/refs/heads/main/install.sh | sh

set -e

REPO="MehrFunds/mehr"
CHAIN_ID="mehr-1"
BINARY="mehrd"
NODE_HOME="${HOME}/.mehr"
REST_URL="http://localhost:1317"
SERVICE_NAME="mehrd"
LOG_FILE="${NODE_HOME}/mehrd.log"
PID_FILE="/tmp/mehrd.pid"

# ── platform detection ────────────────────────────────────────────────────────

detect_platform() {
  OS=$(uname -s | tr '[:upper:]' '[:lower:]')
  ARCH=$(uname -m)

  case "$OS" in
    linux)  ;;
    darwin) ;;
    msys*|mingw*|cygwin*) OS="windows" ;;
    *)
      die "Unsupported OS: $OS — download manually from https://github.com/${REPO}/releases"
      ;;
  esac

  case "$ARCH" in
    x86_64)        ARCH="amd64" ;;
    arm64|aarch64) ARCH="arm64" ;;
    armv7l|armv6l) ARCH="arm"   ;;
    *)
      die "Unsupported architecture: $ARCH"
      ;;
  esac

  EXT=""
  [ "$OS" = "windows" ] && EXT=".exe"
}

# ── helpers ───────────────────────────────────────────────────────────────────

die()  { printf "Error: %s\n" "$1" >&2; exit 1; }
say()  { printf "%s\n" "$1"; }
ask()  { printf "%s " "$1"; read -r REPLY; }

http_get() {
  url="$1"; dest="$2"
  if command -v curl > /dev/null 2>&1; then
    curl -fsSL --progress-bar "$url" -o "$dest"
  elif command -v wget > /dev/null 2>&1; then
    wget -q --show-progress -O "$dest" "$url"
  else
    die "curl or wget is required"
  fi
}

mehrd_bin() {
  command -v "$BINARY" 2>/dev/null || printf "%s" "${BIN_DIR:-/usr/local/bin}/${BINARY}${EXT:-}"
}

# ── install binary ────────────────────────────────────────────────────────────

install_binary() {
  detect_platform
  ASSET="${BINARY}-${OS}-${ARCH}${EXT}"
  URL="https://github.com/${REPO}/releases/latest/download/${ASSET}"

  if [ "$OS" = "windows" ]; then
    BIN_DIR="${HOME}/.local/bin"
    mkdir -p "$BIN_DIR"
  elif [ -w /usr/local/bin ]; then
    BIN_DIR="/usr/local/bin"
  elif [ -w /usr/bin ]; then
    BIN_DIR="/usr/bin"
  else
    BIN_DIR="${HOME}/.local/bin"
    mkdir -p "$BIN_DIR"
  fi

  say "Downloading ${ASSET}..."
  http_get "$URL" "${BIN_DIR}/${BINARY}${EXT}"
  [ "$OS" != "windows" ] && chmod +x "${BIN_DIR}/${BINARY}${EXT}"
  say "Installed ${BIN_DIR}/${BINARY}${EXT}"

  case ":${PATH}:" in
    *":${BIN_DIR}:"*) ;;
    *) say "  Add ${BIN_DIR} to your PATH: export PATH=\"\$PATH:${BIN_DIR}\"" ;;
  esac
}

# ── node init ─────────────────────────────────────────────────────────────────

cmd_init() {
  if [ -d "${NODE_HOME}/config" ]; then
    say "Node already initialized at ${NODE_HOME}"
    ask "Re-initialize? This will erase all data. [y/N]:"
    case "$REPLY" in y|Y) rm -rf "$NODE_HOME" ;; *) return ;; esac
  fi

  ask "Moniker (display name for your node):"; MONIKER="${REPLY:-my-node}"

  "$(mehrd_bin)" init "$MONIKER" --chain-id "$CHAIN_ID"
  say ""
  say "Node initialized. Config is at ${NODE_HOME}/config/"
  say "Edit app.toml and config.toml before starting."
}

# ── systemd service ───────────────────────────────────────────────────────────

has_systemd() {
  command -v systemctl > /dev/null 2>&1
}

service_active() {
  has_systemd && systemctl is-active --quiet "$SERVICE_NAME" 2>/dev/null
}

service_installed() {
  has_systemd && systemctl list-unit-files "${SERVICE_NAME}.service" 2>/dev/null | grep -q "$SERVICE_NAME"
}

cmd_install_service() {
  if ! has_systemd; then
    say "systemd not available — use Start/Stop options to run in background."
    return
  fi
  BIN_PATH=$(mehrd_bin)
  SERVICE_FILE="/etc/systemd/system/${SERVICE_NAME}.service"
  say "Writing ${SERVICE_FILE} (needs sudo)..."
  sudo tee "$SERVICE_FILE" > /dev/null << EOF
[Unit]
Description=Mehr Daemon
After=network-online.target

[Service]
User=${USER}
ExecStart=${BIN_PATH} start
Restart=on-failure
RestartSec=5
LimitNOFILE=65536

[Install]
WantedBy=multi-user.target
EOF
  sudo systemctl daemon-reload
  sudo systemctl enable "$SERVICE_NAME"
  say "Service installed and enabled. Use option 1 to start it."
}

# ── start / stop / status ─────────────────────────────────────────────────────

cmd_start() {
  if service_active; then
    say "mehrd is already running (systemd)."
    return
  fi
  if service_installed; then
    sudo systemctl start "$SERVICE_NAME"
    say "Started via systemd."
    return
  fi
  if [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null; then
    say "mehrd is already running (PID $(cat "$PID_FILE"))."
    return
  fi
  mkdir -p "$(dirname "$LOG_FILE")"
  BIN_PATH=$(mehrd_bin)
  nohup "$BIN_PATH" start >> "$LOG_FILE" 2>&1 &
  printf "%s" "$!" > "$PID_FILE"
  say "Started mehrd (PID $(cat "$PID_FILE")), logging to ${LOG_FILE}."
}

cmd_stop() {
  if service_installed; then
    sudo systemctl stop "$SERVICE_NAME" && say "Stopped via systemd." && return
  fi
  if [ -f "$PID_FILE" ]; then
    PID=$(cat "$PID_FILE")
    kill "$PID" 2>/dev/null && rm -f "$PID_FILE" && say "Stopped mehrd (PID ${PID})." && return
  fi
  say "mehrd is not running."
}

daemon_running() {
  service_active && return 0
  [ -f "$PID_FILE" ] && kill -0 "$(cat "$PID_FILE")" 2>/dev/null
}

# ── logs ──────────────────────────────────────────────────────────────────────

cmd_logs() {
  if service_installed; then
    journalctl -u "$SERVICE_NAME" -n 80 --no-pager 2>/dev/null || true
    return
  fi
  if [ -f "$LOG_FILE" ]; then
    tail -n 80 "$LOG_FILE"
  else
    say "No log file found at ${LOG_FILE}."
  fi
}

# ── status line ───────────────────────────────────────────────────────────────

print_status() {
  if daemon_running; then
    STATUS="running"
    HEIGHT=""
    if command -v curl > /dev/null 2>&1; then
      RAW=$(curl -sf --max-time 2 "${REST_URL}/cosmos/base/tendermint/v1beta1/blocks/latest" 2>/dev/null) || true
      HEIGHT=$(printf "%s" "$RAW" | grep -o '"height":"[^"]*"' | head -1 | cut -d'"' -f4)
    fi
    if [ -n "$HEIGHT" ]; then
      printf "  Status : running   Height : #%s\n" "$HEIGHT"
    else
      printf "  Status : running\n"
    fi
  else
    printf "  Status : stopped\n"
  fi
}

# ── admin panel ───────────────────────────────────────────────────────────────

panel() {
  while true; do
    printf "\n============================================\n"
    printf "  Mehr Daemon Admin\n"
    printf "============================================\n"
    print_status
    printf "--------------------------------------------\n"
    printf "  1) Start daemon\n"
    printf "  2) Stop daemon\n"
    printf "  3) View logs\n"
    printf "  4) Initialize node\n"
    printf "  5) Install as systemd service\n"
    printf "  6) Update binary\n"
    printf "  0) Exit\n"
    printf "--------------------------------------------\n"
    ask "Choice:"

    case "$REPLY" in
      1) cmd_start ;;
      2) cmd_stop ;;
      3) cmd_logs ;;
      4) cmd_init ;;
      5) cmd_install_service ;;
      6) install_binary ;;
      0) exit 0 ;;
      *) say "Invalid choice." ;;
    esac

    printf "\nPress Enter to continue..."; read -r _
  done
}

# ── main ──────────────────────────────────────────────────────────────────────

main() {
  # Skip install if already available and up to date
  if ! command -v "$BINARY" > /dev/null 2>&1; then
    install_binary
  else
    say "Found $(mehrd_bin) — skipping download. Use option 6 to update."
  fi

  if [ ! -d "${NODE_HOME}/config" ]; then
    say ""
    ask "No node found at ${NODE_HOME}. Initialize now? [Y/n]:"
    case "$REPLY" in n|N) ;; *) cmd_init ;; esac
  fi

  panel
}

main "$@"
