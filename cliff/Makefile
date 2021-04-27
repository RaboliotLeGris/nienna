.PHONY: run build test lint fmt

run:
	go run main.go

build: fmt
	go build -o build/cliff

test: lint
	echo "TODO"

lint: fmt
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt -s -w .