#!/bin/bash

#
# This file is auto generated.  Do not edit.
#
# It was created from the JSON config file and shows the env var that can be used to config the app.
# The docker run command will set the env vars on the container.
# You will need to adjust the image name in the Docker command.
#
# The values shown for the env var are the app defaults from the JSON file.
#
# database host name.
# QUAKESEARCH_DATABASE_HOST=localhost
#
# database User password (unencrypted).
# QUAKESEARCH_DATABASE_PASSWORD=test
#
# usually disable or require.
# QUAKESEARCH_DATABASE_SSL_MODE=disable
#
# database connection pool.
# QUAKESEARCH_DATABASE_MAX_OPEN_CONNS=30
#
# database connection pool.
# QUAKESEARCH_DATABASE_MAX_IDLE_CONNS=20
#
# web server port.
# QUAKESEARCH_WEB_SERVER_PORT=8080
#
# public CNAME for the service.
# QUAKESEARCH_WEB_SERVER_CNAME=localhost
#
# true if the app is production.
# QUAKESEARCH_WEB_SERVER_PRODUCTION=false

docker run -e "QUAKESEARCH_DATABASE_HOST=localhost" -e "QUAKESEARCH_DATABASE_PASSWORD=test" -e "QUAKESEARCH_DATABASE_SSL_MODE=disable" -e "QUAKESEARCH_DATABASE_MAX_OPEN_CONNS=30" -e "QUAKESEARCH_DATABASE_MAX_IDLE_CONNS=20" -e "QUAKESEARCH_WEB_SERVER_PORT=8080" -e "QUAKESEARCH_WEB_SERVER_CNAME=localhost" -e "QUAKESEARCH_WEB_SERVER_PRODUCTION=false" busybox
