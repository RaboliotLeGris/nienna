FROM debian:11.2-slim@sha256:d5cd7e54530a8523168473a2dcc30215f2c863bfa71e09f77f58a085c419155b

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.32.1-4+deb11u2 libjson-perl=4.03000-1 libdigest-sha-perl=6.02-1+b3 libclass-dbi-perl=3.0.17-4.1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna