FROM node:17.1.0-alpine@sha256:707d4ffcaafe2bcc9f279fcd40591e0f6eb997be7f1f886796da2c17af871124 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.4.6-alpine

COPY --from=builder /services/webapp/dist/ /usr/share/caddy/