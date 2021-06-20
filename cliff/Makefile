.PHONY: run build test lint fmt

run:
	go run main.go

build:
	go build -o build/cliff

build_image:
	docker build -t cliff:latest .

test: lint

lint: ci-fmt
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

ci-fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt --ci .

fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt -s -w .

test_functional:
	(cd tests && make test)