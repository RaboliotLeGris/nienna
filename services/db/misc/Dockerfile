FROM debian:10.8-slim@sha256:ca9003e9899b458d46b6ee6b040cf9e0d715acbc58e0615871f019c38ada8bc1

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.28.1-6+deb10u1 libjson-perl=4.02000-1 libdigest-sha-perl=6.02-1+b1 libclass-dbi-perl=3.0.17-4 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna