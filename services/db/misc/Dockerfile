FROM debian:10.11-slim@sha256:cfadaa70a452076a478dcfcf2da3a3dc3b741450afed5ab67a7f3df161387c48

RUN apt-get update && apt-get install --no-install-recommends -y perl=5.28.1-6+deb10u1 libjson-perl=4.02000-1 libdigest-sha-perl=6.02-1+b1 libclass-dbi-perl=3.0.17-4 && \
    rm -rf /var/lib/apt/lists/*
WORKDIR /nienna