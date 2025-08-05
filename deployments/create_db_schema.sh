#!/bin/bash

# This script creates a list of PostgreSQL databases and a 'dbo' schema
# inside each, handling cases where they might already exist.
#
# IMPORTANT: This script uses the 'admin' user and password you provided.

# --- Configuration ---
DB_USER="admin"
DB_PASSWORD="admin"
DB_HOST="localhost"
DB_LIST="receivingdb inventorydb fulfillmentdb"

# Export the password as an environment variable for psql and createdb.
export PGPASSWORD="$DB_PASSWORD"

# --- Database Creation ---
echo "--- Creating databases ---"
for db in $DB_LIST; do
  # A more robust way to check for a database's existence by querying the
  # pg_database system catalog, which avoids version-specific psql issues.
  if psql -U "$DB_USER" -h "$DB_HOST" -d "postgres" -c "SELECT 1 FROM pg_database WHERE datname='$db'" -q -t | grep -q '1'; then
    echo "  -> Database '$db' already exists. Skipping creation."
  else
    echo "  -> Creating database '$db'..."
    createdb -U "$DB_USER" -h "$DB_HOST" "$db"
    if [ $? -eq 0 ]; then
      echo "     - Database '$db' created successfully."
    else
      echo "     - Error creating database '$db'."
      # Exit if a database creation fails.
      exit 1
    fi
  fi
done

# --- Schema Creation ---
echo ""
echo "--- Creating 'dbo' schema in each database ---"
for db in $DB_LIST; do
  echo "  -> Creating schema 'dbo' in database '$db'..."
  # Connect to the database and create the schema if it doesn't exist.
  psql -U "$DB_USER" -h "$DB_HOST" -d "$db" -c "CREATE SCHEMA IF NOT EXISTS dbo;"
  if [ $? -eq 0 ]; then
    echo "     - Schema 'dbo' created successfully in '$db'."
  else
    echo "     - Error creating schema 'dbo' in '$db'."
    # Exit if a schema creation fails.
    exit 1
  fi
done

# Unset the PGPASSWORD variable for security after the script is done.
unset PGPASSWORD

echo ""
echo "--- Script finished successfully ---"
