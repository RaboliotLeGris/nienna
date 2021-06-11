FROM mcr.microsoft.com/dotnet/sdk:5.0.6-buster-slim-amd64@sha256:21ed969a93d9fc7c0f86ddc0db7dc72956cb7fd5b5e90bc39fbade6afd01bde9

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]