#!/usr/bin/env bash

CLIFF_HOST="${CLIFF_HOST:-"http://localhost"}"
CLIFF_ROUTE="/api/health"
TRIES=0

while [ `curl -s -o /dev/null -w "%{http_code}\n" $CLIFF_HOST$CLIFF_ROUTE` -ne 200 ] && [ $TRIES -ne 8 ]
do
  TRIES=$((TRIES+1))
  echo "Waiting for Cliff to start up. Tries: $TRIES/8"
  sleep 5
done

if [ $TRIES -ge 8 ]
then
  echo "Failed to access cliff"
  exit 1
fi
