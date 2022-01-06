FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:fedc10704a804fb3ded5bc4a81a82e102eb1ec2a7b219062cc9a08b5566c177f

RUN apt-get update && apt-get install --no-install-recommends -y make=4.2.1-1.2 && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /pulsar

CMD ["bash"]