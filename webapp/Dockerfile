FROM node:15.14.0-alpine@sha256:0944bcebe7fb69f2e81080b879d68e93446d5118fb857f029c3516df25d374d9 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.3.0-alpine

COPY --from=builder /webapp/dist/ /usr/share/caddy/