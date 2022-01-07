FROM mcr.microsoft.com/dotnet/runtime:5.0.13-buster-slim-amd64@sha256:6e7f7cfb89f5f23c2a731905f450fecc73dfebf53fe892dd65e902d6897bca1f AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.404-buster-slim-amd64@sha256:17e8a710a49a73bba03fbe4005fb791851aef8cf2af3efa4ec78767557e410e5 AS build
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