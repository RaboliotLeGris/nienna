FROM node:16 as builder

WORKDIR /webapp
COPY . .
RUN npm i && npm run build

FROM nginx:1.21.4@sha256:366e9f1ddebdb844044c2fafd13b75271a9f620819370f8971220c2b330a9254

COPY --from=builder /webapp/dist/ /usr/share/nginx/html