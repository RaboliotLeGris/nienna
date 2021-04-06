# Nienna

## Services:

![Docs](docs/archi_schema.png)

### Custom
* webapp (vue) -> Allow to upload, watch video and watch livestream
* api (Go) -> serve webapp, and deliver content
* backburner (rust) -> Process uploaded video into hls
* river -> handle incoming livestream and save them

### Tiers
* Bucket storage (minio) -> https://hub.docker.com/r/minio/minio/
* async message bus (Rabbitmq) -> https://hub.docker.com/_/rabbitmq
* Relational database (pgsql) -> https://hub.docker.com/_/postgres
* Reverse proxy (caddy) -> 

## TODOLIST

* backburner -> Do standalone test/package/crates to work on the video processing bits
* backburner -> fetch video from s3
* backburner -> check mimetype
* backburner -> convert it to DASH/HLS
* backburner -> upload them to minio
* backburner -> send event with video status
* backburner -> check worker status
* backburner -> make workers uncrashable .... HF
* cliff -> Do view

* Docker -> makefile with hadolint + on other projects
* cliff -> fix memory usage when uploading file to minio (another client or some tweaks)
* cliff -> Lock database when initializing database
* cliff -> Do not crash when db is not ready
* cliff -> Use resumable ?
* cliff -> Test Cliff
* cliff -> Prevent SQL injection (OMG)
* cliff -> add password

## Interesting articles
* https://blog.eleven-labs.com/fr/video-live-dash-hls/