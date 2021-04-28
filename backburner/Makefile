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
	RUST_LOG=DEBUG cargo test -- --nocapture