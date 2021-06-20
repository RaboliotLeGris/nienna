run:
	go run main.go
build_image:
	docker build -t db:latest .
ci-fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt --ci .
lint: ci-fmt
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v
test: lint

# Perl part
build_perl_image:
	docker build misc -t nienna_perl
update: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl misc/update_schema.pl"
test_schema: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl misc/check_schema.pl"