.PHONY: install build

install:
	npm i

build: install
	npm run build

build-image:
	rm -rf node_module
	docker build .

test: build-image