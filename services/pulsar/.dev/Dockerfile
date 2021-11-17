FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:e2f0b34b64c845636e971449c514cc9f8fffcb486c86cf023baf648040f6cbbf

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]