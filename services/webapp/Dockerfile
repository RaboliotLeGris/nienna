FROM node:16.13.0-alpine@sha256:7001ff8aabeabe8844ef8e7613afaf723e832e19fba55bf1b9929739c7723ece as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM caddy:2.3.0-alpine

COPY --from=builder /services/webapp/dist/ /usr/share/caddy/