FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.32.1-4+deb11u2 libjson-perl=4.03000-1 libdigest-sha-perl=6.02-1+b3 libclass-dbi-perl=3.0.17-4.1 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna