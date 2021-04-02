# Nienna

### Services:

![Docs](docs/archi_schema.png)

#### Custom
* webapp (vue) -> Allow to upload, watch video and watch livestream
* api (Go) -> serve webapp, and deliver content
* backburner (rust) -> Process uploaded video into dash/hls
* river -> handle incoming livestream and save them

#### Tiers
* Bucket storage (minio) -> https://hub.docker.com/r/minio/minio/
* async message bus (Rabbitmq) -> https://hub.docker.com/_/rabbitmq
* Relational database (pgsql) -> https://hub.docker.com/_/postgres
* Reverse proxy (caddy) -> https://hub.docker.com/_/traefik

### TODOLIST

* backfurnace -> worker pool
* backfurnace -> fetch video from s3
* backfurnace -> check mimetype
* backfurnace -> listen to rabbitmq events
* backfurnace -> convert it to DASH/HLS
* backfurnace -> upload them to minio
* backfurnace -> send event with video status
* cliff -> Add status route
* cliff -> Do view

* Docker -> makefile with hadolint + on other projects
* cliff -> fix memory usage when uploading file to minio (another client or some tweaks)
* cliff -> Lock database when initializing database
* cliff -> Do not crash when db is not ready
* cliff -> Use resumable ?
* cliff -> Test Cliff
* cliff -> Prevent SQL injection (OMG)
* cliff -> add password