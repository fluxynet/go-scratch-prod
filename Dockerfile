FROM golang:1.19-bullseye AS builder

COPY . /app
WORKDIR /app

RUN CGO_ENABLED=1 go build -ldflags '-s -w -extldflags "-static"' -o /app/build/serve /app/cmd/serve

FROM debian:bullseye-slim

COPY --from=builder /app/build/ /var/app/

WORKDIR /var/app

EXPOSE 9000
CMD ["./serve"]