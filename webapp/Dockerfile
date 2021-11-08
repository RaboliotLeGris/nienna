FROM node:15.14.0-alpine@sha256:6edd37368174c15d4cc59395ca2643be8e2a1c9846714bc92c5f5c5a92fb8929 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.3.0-alpine

COPY --from=builder /webapp/dist/ /usr/share/caddy/