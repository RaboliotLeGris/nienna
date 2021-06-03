FROM mcr.microsoft.com/dotnet/runtime:5.0.6-buster-slim-amd64@sha256:ca9c27f783381678c8a06d93944469c8cfc6669d7a244c26dd2cbbedfbcfd93f AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.6-buster-slim-amd64@sha256:21ed969a93d9fc7c0f86ddc0db7dc72956cb7fd5b5e90bc39fbade6afd01bde9 AS build
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