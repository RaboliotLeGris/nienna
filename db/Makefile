run:
	go run main.go

build_image:
	docker build -t cliff:latest .

ci-fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt --ci .

build_perl_image:
	docker build misc -t nienna_perl

update: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl misc/update_schema.pl"

test: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl misc/check_schema.pl"