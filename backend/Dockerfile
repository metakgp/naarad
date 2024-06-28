FROM golang:1.22.4 AS builder

WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux go build -o /app --tags "fts5" -a -ldflags '-linkmode external -extldflags "-static"' .

FROM busybox:1.36-musl

COPY --from=builder /app /app
CMD ["/app"]