FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.6@sha256:e1211ac17b29b585ed1aee166a17fad63d344bc973bc63849d74c6452d549b3e

COPY --from=builder /webapp/dist/ /usr/share/nginx/html