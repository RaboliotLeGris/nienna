test: lint
	echo "TODO"

lint: 
	docker run --rm -v $(PWD):/app -w /app golangci/golangci-lint:v1.38.0 golangci-lint run -v

fmt:
	docker run --rm -v $(PWD):/data cytopia/gofmt -s -w .