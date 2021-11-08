#!/usr/bin/env bash

CLIFF_HOST="${CLIFF_HOST:-"http://localhost"}"
CLIFF_ROUTE="/api/health"
TRIES=0

echo "$CLIFF_HOST$CLIFF_ROUTE"
while [ `curl -v -s -o /dev/null -w "%{http_code}\n" $CLIFF_HOST$CLIFF_ROUTE` -ne 200 ] && [ $TRIES -ne 10 ]
do
  docker logs -f "name=nienna_redis"
  echo "`curl -v -s -o /dev/null -w "%{http_code}\n" $CLIFF_HOST$CLIFF_ROUTE`"
  TRIES=$((TRIES+1))
  echo "Waiting for Cliff to start up. Tries: $TRIES/8"
  sleep 5
done
  echo "`curl -v -s -o /dev/null -w "%{http_code}\n" $CLIFF_HOST$CLIFF_ROUTE`"

if [ $TRIES -ge 10 ]
then
  echo "Failed to access cliff"
  exit 1
fi
