.PHONY: run test

start_env:
	echo "Start testing env"
	docker-compose -p nienna up -d --build --remove-orphans
	chmod +x wait_for_cliff.sh && ./wait_for_cliff.sh

wait_for_cliff:
	docker run --rm --network nienna_default --env CLIFF_HOST="http://cliff:8000" -v $$PWD/wait_for_cliff.sh:/tmp/wait_for_cliff.sh debian /tmp/wait_for_cliff.sh

stop_env:
	echo "Stopping testing env"
	docker stop $$(docker ps -q)
	# TODO: DELETE ALL VOLUMES

exec_test:
	go test -v ./... -count=1

run:
	echo "RUNNING functional tests"
	docker build . -t nienna-functional
	docker run --rm --network nienna_default --env DB_URI="postgresql://db/nienna?user=nienna&password=nienna" -v $$PWD/../samples:/go/src/cliff/samples --env CLIFF_HOST="http://cliff:8000" nienna-functional 

test:
	make -s start_env
	make -k run
	make -k stop_env
	