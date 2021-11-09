FROM mcr.microsoft.com/dotnet/runtime:5.0.12-buster-slim-amd64@sha256:1c0b93fac568d1fe3d81a048758a02985b3ec55884a5f8b3557ba6a1be7e0d78 AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:af2f5f5ba4109923d22276e8852249047bd462f22669e48184f59dea22e26c46 AS build
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