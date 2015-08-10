#!/bin/bash

# script for initializing the db in the postgres Docker container.

export PGUSER=postgres

cd /docker-entrypoint-initdb.d

psql  -d postgres < /docker-entrypoint-initdb.d/create-users.ddl
psql  -d postgres < /docker-entrypoint-initdb.d/create-db.ddl
psql  -d hazard -c 'create extension postgis;'
psql  --quiet hazard < /docker-entrypoint-initdb.d/wfs.ddl
psql  --quiet hazard < /docker-entrypoint-initdb.d/wfs-region.ddl
psql  --quiet hazard < /docker-entrypoint-initdb.d/wfs-region-values.ddl
psql  --quiet hazard < /docker-entrypoint-initdb.d/wfs-event-values.ddl
psql  --quiet hazard < /docker-entrypoint-initdb.d/user-permissions.ddl
