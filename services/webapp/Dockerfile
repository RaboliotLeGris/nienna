FROM node:16.13.0-bullseye@sha256:c37d94a3225dd200a75912eba2bae0692affb5bb16818b9d1d947896706e8abc as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.4.6@sha256:7da0f90273e1961d9c38d26809f84d4ef3cdc9b4fc330a9cab22015d7c9e8228

COPY --from=builder /webapp/dist/ /usr/share/caddy/