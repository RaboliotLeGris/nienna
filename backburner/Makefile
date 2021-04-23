.PHONY: run build build-release test

run:
	cargo run

build:
	cargo build

build-release:
	cargo build --release

test:
	RUST_LOG=DEBUG cargo test -- --nocapture

d:
	docker build -t dev-backburner .dev && docker run -ti --rm --env RUST_LOG=DEBUG -v $$PWD:/usr/src/backburner dev-backburner