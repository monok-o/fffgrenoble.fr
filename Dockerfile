# Building frontend
FROM golang:1.17-alpine AS builder
WORKDIR /build
RUN apk --no-cache add ca-certificates
RUN apk add --update nodejs npm
COPY . .
RUN npm install
RUN npm run build
RUN go mod tidy
RUN go build main.go

# running app
FROM alpine:latest  
WORKDIR /app
COPY --from=builder /build/view/public /app/view/public
COPY --from=builder /build/main /app/main
CMD ["/app/main"]
