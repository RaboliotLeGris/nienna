.PHONY: run build test

run:
	(cd pulsar && dotnet run)

build_image:
	docker build -t pulsar .

build_dev:
	docker build .dev -t nienna_pulsar_test

test: build_dev
	docker run --rm -v $$PWD:/pulsar --env RUST_LOG="DEBUG" nienna_pulsar_test dotnet test


