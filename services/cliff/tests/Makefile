.PHONY: run test

test_start_env:
	echo "Start testing env"
	docker-compose -p nienna up -d --build --remove-orphans
	chmod +x wait_for_cliff.sh && ./wait_for_cliff.sh

test_stop_env:
	echo "Stopping testing env"
	docker stop $$(docker ps -q)
	docker container rm $$(docker container ls -aq)
	docker volume rm -f $$(docker volume ls -q)

test_exec:
	go test -timeout 1800s -v ./... -count=1

test_run:
	echo "RUNNING functional tests"
	(cd .. && docker build tests -t nienna-functional)
	docker run --rm --network nienna_default --env DB_URI="postgresql://pg/nienna?user=nienna&password=nienna" -v $$PWD/../../../samples:/go/src/cliff/samples --env CLIFF_HOST="http://cliff:8000" nienna-functional

test:
	make -s test_start_env
	make -k test_run
	make -k test_stop_env
	