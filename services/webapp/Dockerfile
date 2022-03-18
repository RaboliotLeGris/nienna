FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.6@sha256:9e3049eff104b703e08b67cc5c0c2e0ad1973d4b9185d8c164b7958bd283dbf3

COPY --from=builder /webapp/dist/ /usr/share/nginx/html