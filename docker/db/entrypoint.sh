#!/bin/bash

# Check if the environment variables are set
if [ -z "$MYSQL_ROOT_PASSWORD" ]; then
    echo "Error: MYSQL_ROOT_PASSWORD environment variable is not set."
    exit 1
fi

if [ -z "$MYSQL_USER" ]; then
    echo "Error: MYSQL_USER environment variable is not set."
    exit 1
fi

if [ -z "$MYSQL_PASSWORD" ]; then
    echo "Error: MYSQL_PASSWORD environment variable is not set."
    exit 1
fi

if [ -z "$MYSQL_DATABASE" ]; then
    echo "Error: MYSQL_DATABASE environment variable is not set."
    exit 1
fi

# Create the SQL query for creating the database
CREATE_DB_QUERY="CREATE DATABASE IF NOT EXISTS \`$MYSQL_DATABASE\`;"

# Run the SQL query
echo "Creating database '$MYSQL_DATABASE'..."
mysql -u root -p"$MYSQL_ROOT_PASSWORD" -e "$CREATE_DB_QUERY"
echo "Database '$MYSQL_DATABASE' created."

# Run the init sql file
echo "Running the init_db.sql file..."
mysql -u $MYSQL_USER -p"$MYSQL_PASSWORD" $MYSQL_DATABASE < /init_sql/insert_app_version_default_record.sql
echo "init_db.sql executed."