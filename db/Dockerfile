FROM debian:10.8-slim@sha256:13f0764262a064b2dd9f8a828bbaab29bdb1a1a0ac6adc8610a0a5f37e514955

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.28.1-6+deb10u1 libjson-perl=4.02000-1 libdigest-sha-perl=6.02-1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna