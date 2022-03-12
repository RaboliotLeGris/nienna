.PHONY: run build test lint fmt

run:
	go run main.go

build:
	go build -o build/river

build_image:
	docker build -t river:latest .

test:
	go test ./...

lint: ci-fmt
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

ci-fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt --ci .

fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt -s -w .
