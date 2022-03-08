FROM mcr.microsoft.com/dotnet/runtime:5.0.15-buster-slim-amd64@sha256:95bd0886e51e8d8720151e219f09065fbe7338349d78478a7d8948f637156aaf AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.404-buster-slim-amd64@sha256:75ff3ad3a567482d7403eeeef5424fac254c5931a57992e07c2513204168c1a4 AS build
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