# Nienna

Nienna is a solution to upload and watch videos. It converts uploaded videos to HLS and serve it on demand.

It later aim to allow streaming.

## Name origin:

"Tutor of Olórin; weeps constantly, but not for herself; and those who hearken to her learn pity, and endurance in hope.
She gives strength to those in the Hall of Mandos. Her tears are those of healing and compassion, not of sadness, and
often have potency; she watered the Two Trees with her tears, and washed the filth of Ungoliant away from them once they
were destroyed. She was in favour of releasing Melkor after his sentence, not being able to see his evil nature."
*Source: [Wikipedia](https://en.wikipedia.org/wiki/Vala_(Middle-earth)#Nienna)*

## How to use
To run it locally, the easiest way is to use docker and execute `make d_launch`. 

## How to develop
It can be difficult, specially if you use Windows. You will need to have installed on your system `Node`, `Go`, `ffmpeg`, `Rust`, `Python3`, `RabbitMQ`, `PostgreSQL` and `Minio`. You might also run into some troubles with CORS.

Or you can run: `make d_dev` in one terminal and while in `webapp` dir, `npm run serve`. 

It starts all the containers but it also mount each services without launching them. This way you can hop into them with `make d_<SERVICE_NAME>` and either launch it while needing it without have to relaunch all the services.

## Services:

![Docs](docs/archi_schema.png)

### Custom

* Webapp (Vue) -> Allow to upload, watch video and watch livestream
* Cliff (Go) -> serve webapp, and serve http request (handle uploaded file and watching)
* Backburner (Rust) -> Process uploaded video into HLS
* Pulsar (Python) -> Save event from videos
* River -> handle incoming livestream and save them (done later)

### Tiers

* Bucket storage (Minio) -> https://hub.docker.com/r/minio/minio/
* Asynchronous message bus (RabbitMQ) -> https://hub.docker.com/_/rabbitmq
* Relational database (Pgsql) -> https://hub.docker.com/_/postgres
* Reverse proxy (Caddy) -> https://hub.docker.com/_/caddy

## TODOLIST
* Run tests on github CI
* cliff -> Test Cliff
* cliff -> fix return of struct to remove first letter with Uppercase
* Doc -> flow
* Backburner -> check worker status -> Heartbeat (not doable yet)
* Docker -> makefile with hadolint + on other projects
* cliff -> fix memory usage when uploading file to minio (another client or some tweaks)
* cliff -> Lock database when initializing database
* cliff -> Do not crash when db is not ready
* cliff -> Use resumable ?
* cliff -> Prevent SQL injection (OMG)
* cliff -> add password
