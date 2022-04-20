FROM node:17 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.6@sha256:2275af0f20d71b293916f1958f8497f987b8d8fd8113df54635f2a5915002bf1

COPY --from=builder /webapp/dist/ /usr/share/nginx/html