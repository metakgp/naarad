FROM golang:1.22.4 AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o /app --tags "fts5" -a -ldflags '-linkmode external -extldflags "-static"' .

FROM busybox:1.36-musl

RUN mkdir -p /etc/ssl/certs

COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

COPY metaploy/naarad-api.metaploy.conf /

COPY metaploy/postinstall.sh /

RUN chmod +x /postinstall.sh

COPY --from=builder /app /app

CMD ["/postinstall.sh", "/app"]