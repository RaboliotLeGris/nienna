
.PHONY: build test d_launch

d_launch:
	docker-compose up --build --remove-orphans

build: build_api build_backburner build_webapp

test: test_api test_backburner test_webapp

build_webapp:
	(cd webapp && npm i && npm run build)
	(rm -rf webapi/static/*; cp -r webapp/dist/* webapi/static/)
build_backburner:
	(cd backburner && cargo build)
build_webapi:
	(cd webapi && go build -o build/webapi)

test_webapp:
	echo 'TODO'
test_backburner:
	(cd backburner && cargo test)
test_webapi:
	echo 'TODO'

