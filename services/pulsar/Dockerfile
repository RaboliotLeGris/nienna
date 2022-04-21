FROM mcr.microsoft.com/dotnet/runtime:5.0.16-buster-slim-amd64@sha256:d21ae89d5d13a39888acb0258a32face4eba76f9a7df2cc1dc9f81ec5d0e0011 AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.407-buster-slim-amd64@sha256:a884882a177081a38efe2cc8aa6f38243a716c7466e33296e5dfb336fe46f544 AS build
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