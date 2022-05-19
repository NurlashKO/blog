FROM golang:1.18-buster as builder

ENV CGO_ENABLED 0

WORKDIR /opt/blog

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /app ./...


FROM debian:buster-slim

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

COPY --from=builder /app /app

CMD ["/app"]