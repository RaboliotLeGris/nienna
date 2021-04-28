.PHONY: run test

start_env:
	echo "Start testing env"
	docker-compose -p nienna up -d --build --remove-orphans
	./wait_for_cliff.sh

stop_env:
	echo "Stopping testing env"
	docker stop $$(docker ps -q)
	# TODO: DELETE ALL VOLUMES

run:
	echo "RUNNING functional tests"
	go test ./... -count=1

test: start_env run #stop_env
	