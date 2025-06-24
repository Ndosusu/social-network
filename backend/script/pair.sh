#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
ENV_FILE="$SCRIPT_DIR/../config/config.env"

echo SCRIPT_DIR: "$SCRIPT_DIR"

if [[ -f "$ENV_FILE" ]]; then
  source "$ENV_FILE"
else
  echo "Error: .env file not found at $ENV_FILE"
  exit 1
fi

migrate create -ext sql -dir "$MIG_PATH" -seq "$@"