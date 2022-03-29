FROM mcr.microsoft.com/dotnet/sdk:5.0.406-buster-slim-amd64@sha256:654ebc0abbdb1b49f446dfe9a12cad3c47867af03a835426d98bec6c3f920057

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]