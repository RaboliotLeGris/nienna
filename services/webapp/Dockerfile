FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.4@sha256:2f14a471f2c2819a3faf88b72f56a0372ff5af4cb42ec45aab00c03ca5c9989f

COPY --from=builder /webapp/dist/ /usr/share/nginx/html