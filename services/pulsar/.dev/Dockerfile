FROM mcr.microsoft.com/dotnet/sdk:5.0.404-buster-slim-amd64@sha256:17e8a710a49a73bba03fbe4005fb791851aef8cf2af3efa4ec78767557e410e5

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]