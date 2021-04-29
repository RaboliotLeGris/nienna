#!/usr/bin/env bash

CLIFF_HOST="http://localhost"
CLIFF_ROUTE="/api/health"
TRIES=0

while [ `curl -sI -o /dev/null -w "%{http_code}\n" $CLIFF_HOST$CLIFF_ROUTE` -ne 200 ] && [ $TRIES -ne 8 ]
do
  TRIES=$((TRIES+1))
  echo "Waiting for Cliff to start up. Tries: $TRIES/8"
  sleep 5
done
