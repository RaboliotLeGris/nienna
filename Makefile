
.PHONY: build test d_launch

d_launch:
	docker-compose -f .docker/docker-compose.app.yml -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans

d_dev:
	docker-compose -f .docker/docker-compose.app.dev.yml -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans

d_launch_services:
	docker-compose -f .docker/docker-compose.services.yml -p nienna up --build --remove-orphans

d_backburner:
	docker exec -ti `docker ps -aqf "name=nienna_backburner"` bash

d_db:
	docker exec -ti `docker ps -aqf "name=nienna_db"` psql --user nienna

d_redis:
	docker exec -ti `docker ps -aqf "name=nienna_redis"` redis-cli

build: build_cliff build_backburner build_webapp

test: test_cliff test_backburner test_webapp

build_webapp:
	(cd webapp && make build)
	(rm -rf cliff/static/*; cp -r webapp/dist/* cliff/static/)
build_backburner:
	(cd backburner && make build)
build_cliff:
	(cd cliff && make build)

test_webapp:
	echo 'TODO'
test_backburner:
	echo 'TODO'
test_cliff:
	(cd cliff && make test)

