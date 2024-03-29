.PHONY: build test d_launch

d_launch: clean
	docker-compose -f .docker/docker-compose.app.yml -f .docker/docker-compose.services.yml -p nienna up -d --build --remove-orphans
d_dev: clean
	docker-compose -f .docker/docker-compose.app.dev.yml -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans

d_launch_services:
	docker-compose -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans
d_backburner:
	docker exec -ti `docker ps -aqf "name=nienna_backburner"` bash
d_cliff:
	docker exec -ti `docker ps -aqf "name=nienna_cliff"` bash
d_pulsar:
	docker exec -ti `docker ps -aqf "name=nienna_pulsar"` bash
d_pg:
	docker exec -ti `docker ps -aqf "name=nienna_pg"` psql --user nienna
d_redis:
	docker exec -ti `docker ps -aqf "name=nienna_redis"` redis-cli

clean:
	rm -rf backburner/target

build: build_cliff build_backburner build_webapp
test: test_cliff test_backburner test_pulsar test_functional test_dockerfiles

build_images: build_webapp
	(cd backburner && make build_image)
	(cd cliff && make build_image)
	(cd pulsar && make build_image)

publish_images:
	go run .deploy/publish_images.go

build_webapp:
	(cd webapp && make build)
build_backburner:
	(cd backburner && make build)
build_cliff:
	(cd cliff && make build)

test_backburner:
	(cd backburner && make test)
test_cliff:
	(cd cliff && make test)
test_pulsar:
	(cd pulsar && make test)
test_functional:
	(cd cliff && make -k test_functional)
test_dockerfiles:
	(cd .docker && chmod +x hadolint.sh && make lint)
test_schema:
	(cd db && make test_schema)
test_db:
	(cd db && make test)
