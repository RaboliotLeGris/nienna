FROM node:16.13.0-alpine@sha256:3bca55259ada636e5fee8f2836aba7fa01fed7afd0652e12773ad44af95868b9 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.4.6-alpine

COPY --from=builder /services/webapp/dist/ /usr/share/caddy/