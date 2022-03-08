FROM mcr.microsoft.com/dotnet/sdk:5.0.404-buster-slim-amd64@sha256:75ff3ad3a567482d7403eeeef5424fac254c5931a57992e07c2513204168c1a4

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]