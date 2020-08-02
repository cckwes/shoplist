#################################
# Stage 1 build binary
#################################
FROM golang:1.14-stretch AS builder

WORKDIR /app
COPY . .

RUN go mod download
RUN GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o shoplist .

#################################
# Stage 2 build a smaller image
#################################
FROM debian:stretch

WORKDIR /app
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /app/shoplist .

EXPOSE 3000

ENTRYPOINT ["/app/shoplist"]
