FROM debian:11.3-slim@sha256:78fd65998de7a59a001d792fe2d3a6d2ea25b6f3f068e5c84881250373577414

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.32.1-4+deb11u2 libjson-perl=4.03000-1 libdigest-sha-perl=6.02-1+b3 libclass-dbi-perl=3.0.17-4.1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna