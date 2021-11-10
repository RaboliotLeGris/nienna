FROM debian:10.11-slim@sha256:544c93597c784cf68dbe492ef35c00de7f4f6a990955c7144a40b20d86a3475f

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.28.1-6+deb10u1 libjson-perl=4.02000-1 libdigest-sha-perl=6.02-1+b1 libclass-dbi-perl=3.0.17-4 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna