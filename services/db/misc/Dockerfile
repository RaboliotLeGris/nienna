FROM debian:11.1-slim@sha256:a23887a2e830b815955e010f30d4c2430cd5ef82e93c130471024bc9f808d5d3

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.28.1-6+deb10u1 libjson-perl=4.02000-1 libdigest-sha-perl=6.02-1+b1 libclass-dbi-perl=3.0.17-4 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna