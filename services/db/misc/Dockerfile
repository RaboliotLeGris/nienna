FROM debian:11.2-slim@sha256:b0d53c872fd640c2af2608ba1e693cfc7dedea30abcd8f584b23d583ec6dadc7

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.32.1-4+deb11u2 libjson-perl=4.03000-1 libdigest-sha-perl=6.02-1+b3 libclass-dbi-perl=3.0.17-4.1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna