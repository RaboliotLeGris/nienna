.PHONY: build test d_launch

d_launch: clean
	docker-compose -f .docker/docker-compose.app.yml -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans
d_launch_bg: clean
	docker-compose -f .docker/docker-compose.app.yml -f .docker/docker-compose.services.yml -p nienna up -d --remove-orphans
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
	(cd services/backburner && make build_image)
	(cd services/cliff && make build_image)
	(cd services/pulsar && make build_image)

publish_images:
	go run .deploy/publish_images.go

build_webapp:
	(cd services/webapp && make build)
build_backburner:
	(cd services/backburner && make build)
build_cliff:
	(cd services/cliff && make build)

test_backburner:
	(cd services/backburner && make test)
test_cliff:
	(cd services/cliff && make test)
test_pulsar:
	(cd services/pulsar && make test)
test_functional:
	(cd tests/functional && make -k test)
test_dockerfiles:
	(cd .docker && chmod +x hadolint.sh && make lint)
test_schema:
	(cd services/db && make test_schema)
test_db:
	(cd services/db && make test)
test_webapp:
	(cd services/webapp && make test)
