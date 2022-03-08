FROM mcr.microsoft.com/dotnet/runtime:5.0.13-buster-slim-amd64@sha256:72ec4934747e3011504c172ae09dd83cb575fbff3b30dae48a06c281a336af83 AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.406-buster-slim-amd64@sha256:d11f0b2c983ad35ae7b129745ce4bc33c29720f4e746658fb856a776271c013e AS build
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