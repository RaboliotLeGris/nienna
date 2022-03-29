FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.6@sha256:e48e9d28dd773886b5c4b86db4e411eedf46bb98095f5c03c3f6a167a633dcf0

COPY --from=builder /webapp/dist/ /usr/share/nginx/html