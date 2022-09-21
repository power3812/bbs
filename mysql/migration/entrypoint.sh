#!/bin/sh
ls  migrations
/wait
/migrate \
  -path /migrations/ \
  -database "mysql://${MYSQL_USER}:${MYSQL_PASSWORD}@tcp(${MYSQL_HOST}:${MYSQL_PORT})/${MYSQL_DATABASE}" \
  $@
