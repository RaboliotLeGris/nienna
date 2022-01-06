FROM mcr.microsoft.com/dotnet/runtime:5.0.12-buster-slim-amd64@sha256:fc6e9766ee6064a92dce444cfb96b2c8837bb3738d99071a867814eed27970c8 AS base
WORKDIR /pulsar

FROM mcr.microsoft.com/dotnet/sdk:5.0.403-buster-slim-amd64@sha256:fedc10704a804fb3ded5bc4a81a82e102eb1ec2a7b219062cc9a08b5566c177f AS build
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