FROM node:12-alpine AS builder
WORKDIR /building
COPY . .
RUN npm install
RUN npm run build

FROM lukechannings/deno:latest
WORKDIR /app
COPY ./public .
USER deno
CMD ["run", "--allow-net", "--allow-read", "server.js"]
