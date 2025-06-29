#!/bin/bash
SCRIPT_DIR=$(dirname "$0")
ENV_FILE="$SCRIPT_DIR/../config/config.env"

if [[ -f "$ENV_FILE" ]]; then
  source "$ENV_FILE"
else
  echo "Error: .env file not found at $ENV_FILE"
  exit 1
fi

# Calculate the base directory (one level up from script directory)
BASE_DIR="$SCRIPT_DIR/.."

# Create absolute paths
DB_PATH="$BASE_DIR/data/$DB_NAME"
ABSOLUTE_MIG_PATH="$BASE_DIR/pkg/db/migrations"

# Debug: print the paths
echo "Base directory: $BASE_DIR"
echo "Migration path: $ABSOLUTE_MIG_PATH"
echo "Database path: $DB_PATH"


migrate -path $ABSOLUTE_MIG_PATH -database "sqlite3://$DB_PATH" "$@"