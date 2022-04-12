FROM mcr.microsoft.com/dotnet/runtime:5.0.16-buster-slim-amd64@sha256:3cc2269da4a7e40248f39ce08a5e74288d6e59d4be68b57927ce36205847df60 AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.406-buster-slim-amd64@sha256:654ebc0abbdb1b49f446dfe9a12cad3c47867af03a835426d98bec6c3f920057 AS build
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