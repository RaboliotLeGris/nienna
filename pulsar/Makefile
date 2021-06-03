.PHONY: run build test

run:
	(cd pulsar && dotnet run)

build_image:
	docker build -t pulsar .

build_dev:
	docker build .dev -t nienna_pulsar_test

test: build_dev test_setup_env
	docker run --network nienna_default --rm -v $$PWD:/pulsar nienna_pulsar_test dotnet test
	make test_stop_env

test_setup_env:
	docker-compose -f .dev/docker-compose.yml -p nienna up -d --build --remove-orphans
	sleep 30

test_run:
	docker run --network nienna_default --rm -v $$PWD:/pulsar nienna_pulsar_test dotnet test

test_stop_env:
	echo "Stopping testing env"
	docker stop $$(docker ps -q)
	docker container rm $$(docker container ls -aq)
	docker volume rm -f $$(docker volume ls -q)


