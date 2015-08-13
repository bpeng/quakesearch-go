#!/bin/bash

# script for initializing the db in the postgres Docker container.

export PGUSER=postgres

psql -h 127.0.0.1 -d postgres -f  ddl/create-users.ddl
psql -h 127.0.0.1 -d postgres -f ddl/create-db.ddl
psql -h 127.0.0.1 -d hazard -c 'create extension postgis;'
psql -h 127.0.0.1 --quiet hazard -f ddl/wfs.ddl
psql -h 127.0.0.1 --quiet hazard -f ddl/wfs-region.ddl
psql -h 127.0.0.1 --quiet hazard -f ddl/wfs-region-values.ddl
psql -h 127.0.0.1 --quiet hazard -f ddl/wfs-event-values.ddl
psql -h 127.0.0.1 --quiet hazard -f ddl/user-permissions.ddl
