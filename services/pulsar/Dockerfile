FROM mcr.microsoft.com/dotnet/runtime:5.0.12-buster-slim-amd64@sha256:651e23e0bdb4e6b5ac354558924315e2aafd2c0c7705692ae2df1ecd8f6e089d AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:e2f0b34b64c845636e971449c514cc9f8fffcb486c86cf023baf648040f6cbbf AS build
WORKDIR /src
COPY ["pulsar/pulsar.csproj", "pulsar/"]

RUN dotnet restore "pulsar/pulsar.csproj"
COPY . .
WORKDIR "/src/pulsar"
RUN dotnet build "pulsar.csproj" -c Release -o /pulsar/build

FROM build AS publish
RUN dotnet publish "pulsar.csproj" -c Release -o /pulsar/publish

FROM base AS final
WORKDIR /pulsar
COPY --from=publish /pulsar/publish .
ENTRYPOINT ["dotnet", "pulsar.dll"]