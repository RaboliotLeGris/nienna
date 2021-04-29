.PHONY: run build build-release test

run:
	cargo run

build:
	cargo build

build-release:
	cargo build --release

build_image:
	docker build -t backburner .

test:
	docker build .dev -t nienna_backburner_test
	docker run --rm -v $$PWD:/usr/src/backburner --env RUST_LOG="DEBUG" nienna_backburner_test cargo test -- --nocapture
