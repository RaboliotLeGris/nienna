FROM mcr.microsoft.com/dotnet/sdk:5.0.406-buster-slim-amd64@sha256:d11f0b2c983ad35ae7b129745ce4bc33c29720f4e746658fb856a776271c013e

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]