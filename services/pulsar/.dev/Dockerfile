FROM mcr.microsoft.com/dotnet/sdk:5.0.407-buster-slim-amd64@sha256:ad9824d965bb544dd9eeeb31c7525b2d191830137b91dd3a8ef747651104592a

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]