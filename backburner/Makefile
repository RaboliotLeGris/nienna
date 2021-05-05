.PHONY: run build build-release test

run:
	cargo run

build:
	cargo build

build-release:
	cargo build --release

build_image:
	docker build -t backburner .

build_dev:
	docker build .dev -t nienna_backburner_test

test: build_dev clippy
	docker run --rm -v $$PWD:/usr/src/backburner --env RUST_LOG="DEBUG" nienna_backburner_test cargo test -- --nocapture

clippy: build_dev
	docker run --rm -ti -v $$PWD:/usr/src/backburner --env RUST_LOG="DEBUG" nienna_backburner_test cargo clippy