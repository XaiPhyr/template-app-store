#!/bin/bash
set -euo pipefail

#
# $POSTGRES_PASSWORD
#

echo -e "\n"
echo -e "start of migration...\n"

psql -a -v ON_ERROR_STOP=1 -U "$POSTGRES_USER" -d "$POSTGRES_DB" <<-EOSQL
	GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO rdev;

	\connect app_store_db
	CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
EOSQL

echo -e "end of migration..."