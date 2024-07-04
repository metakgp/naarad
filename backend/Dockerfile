FROM golang:1.22.4 AS builder

WORKDIR /src

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=1 GOOS=linux go build -o ./build -a -ldflags '-linkmode external -extldflags "-static"'

FROM alpine:latest AS app

RUN apk --no-cache add ca-certificates tzdata bash

ENV TZ="Asia/Kolkata"

WORKDIR /app

COPY metaploy/ ./

RUN chmod +x ./postinstall.sh

COPY --from=builder /src/build .

CMD ["./postinstall.sh", "./build"]