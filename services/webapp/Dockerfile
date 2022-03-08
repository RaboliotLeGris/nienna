FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.6@sha256:1c13bc6de5dfca749c377974146ac05256791ca2fe1979fc8e8278bf0121d285

COPY --from=builder /webapp/dist/ /usr/share/nginx/html