.PHONY: run

build_image:
	docker build -t pulsar:latest .

run: install
	python3 src/pulsar.py

install:
	pip install -r requirements.txt

freeze:
	pip freeze > requirements.txt
