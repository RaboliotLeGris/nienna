# Nienna Cliff

## Env var config

* NIENNA_DEV -> To determine if Cliff is in dev mode (Extra log + CORS disabled) (default: false)
* DISABLE_NIENNA_REGISTER -> Disallow the creation of new users (default: false)
* DB_URI -> Postgresql db uri
* REDIS_URI -> Redis uri
* REDIS_PASSWORD -> Redis password (optional)
* AMQP_URI -> AMQP uri
* S3_URI -> S3 uri
* S3_ACCESS_KEY -> S3 user key
* S3_SECRET_KEY -> S3 user password
* S3_DISABLE_TLS -> disable https request to S3 (default: false)
* PORT -> application port (default: 8000)