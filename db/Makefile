run:
	perl src/seed.pl



build_perl_image:
	docker build src/misc -t nienna_perl

update: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl src/misc/update_schema.pl"

test: build_perl_image
	docker run --rm -v $(PWD)/..:/nienna nienna_perl sh -c "cd db && perl src/misc/check_schema.pl"