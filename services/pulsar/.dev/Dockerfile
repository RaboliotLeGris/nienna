FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:af2f5f5ba4109923d22276e8852249047bd462f22669e48184f59dea22e26c46

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]