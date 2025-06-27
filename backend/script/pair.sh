#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
ENV_FILE="$SCRIPT_DIR/../config/config.env"

if [[ -f "$ENV_FILE" ]]; then
  source "$ENV_FILE"
else
  echo "Error: .env file not found at $ENV_FILE"
  exit 1
fi

BASE_DIR="$SCRIPT_DIR/../.."
ABSOLUTE_MIG_PATH="$BASE_DIR/$MIG_PATH"

migrate create -ext sql -dir "$ABSOLUTE_MIG_PATH" -seq "$@"