FROM debian:11.1-slim@sha256:656b29915fc8a8c6d870a1247aad7796ce296729b0ae95e168d1cfe30dd2fb3c

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.32.1-4+deb11u2 libjson-perl=4.03000-1 libdigest-sha-perl=6.02-1+b3 libclass-dbi-perl=3.0.17-4.1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna