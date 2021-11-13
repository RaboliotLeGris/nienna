FROM node:17.1.0-alpine@sha256:993bdfb0da7ae8fa4dad7282f797e3e26e88f810d410e0b0409d132d1fb17af3 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.4.6-alpine

COPY --from=builder /services/webapp/dist/ /usr/share/caddy/