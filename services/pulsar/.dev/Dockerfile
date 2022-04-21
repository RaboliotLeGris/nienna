FROM mcr.microsoft.com/dotnet/sdk:5.0.407-buster-slim-amd64@sha256:a884882a177081a38efe2cc8aa6f38243a716c7466e33296e5dfb336fe46f544

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]