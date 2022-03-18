FROM mcr.microsoft.com/dotnet/sdk:5.0.406-buster-slim-amd64@sha256:7f6606602561048c0783e6c77967956bdd980aa5b322a2608f7429f88e3c9008

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]