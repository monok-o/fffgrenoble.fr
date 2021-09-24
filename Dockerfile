# Building frontend
FROM node:12-alpine AS front-builder
WORKDIR /front
COPY ./view /front/view
COPY package.json .
RUN npm install
RUN npm run build

# Building full app
FROM golang:1.17 AS srv-builder
WORKDIR /building
COPY . .
RUN go mod tidy
RUN go build main.go

# running app
FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=front-builder /front/view/public /app/view/public
COPY --from=srv-builder /building/main /app/main
CMD ["/app/main"]
