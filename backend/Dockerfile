
FROM golang:1.22.4 AS builder
WORKDIR /src
COPY . .
RUN go mod download
RUN CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o app
CMD ["./app"]
