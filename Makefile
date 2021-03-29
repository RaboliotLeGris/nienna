
.PHONY: build test d_launch

d_launch:
	docker-compose -f .docker/docker-compose.app.yml -f .docker/docker-compose.services.yml up --build --remove-orphans

d_launch_app:
	docker build cliff -t cliff && docker run --env-file cliff/.env.dev.list --network=docker_default --rm -ti cliff

d_launch_services:
	docker-compose -f .docker/docker-compose.services.yml up --build --remove-orphans

d_db:
	docker exec -ti `docker ps -aqf "name=nienna_db"` psql --user nienna

d_redis:
	docker exec -ti `docker ps -aqf "name=nienna_redis"` redis-cli

build: build_cliff build_backburner build_webapp

test: test_cliff test_backburner test_webapp

build_webapp:
	(cd webapp && npm i && npm run build)
	(rm -rf cliff/static/*; cp -r webapp/dist/* cliff/static/)
build_backburner:
	(cd backburner && cargo build)
build_cliff:
	(cd cliff && go build -o build/cliff)

test_webapp:
	echo 'TODO'
test_backburner:
	echo 'TODO'
test_cliff:
	(cd cliff && pwd && make test)

