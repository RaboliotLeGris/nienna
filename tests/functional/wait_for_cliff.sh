#!/usr/bin/env bash

ENDPOINT_HOST="${ENDPOINT_HOST:-"http://localhost"}"
ROUTE="/api/health"
TRIES=0

echo "$ENDPOINT_HOST$CLIFF_ROUTE"
while [ `curl -v -s -o /dev/null -w "%{http_code}\n" $ENDPOINT_HOST$ROUTE` -ne 200 ] && [ $TRIES -ne 10 ]
do
  docker logs -f "name=nienna_cliff"
  echo "Targeting: $ENDPOINT_HOST$ROUTE"
  echo "`curl -v -s -o /dev/null -w "%{http_code}\n" $ENDPOINT_HOST$ROUTE`"
  TRIES=$((TRIES+1))
  echo "Waiting for Cliff to start up. Tries: $TRIES/10"
  sleep 5
done
  echo "`curl -v -s -o /dev/null -w "%{http_code}\n" $ENDPOINT_HOST$ROUTE`"

if [ $TRIES -ge 10 ]
then
  echo "Failed to access cliff"
  exit 1
fi
