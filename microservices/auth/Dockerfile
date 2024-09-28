FROM golang:1.22.1-bookworm as builder

ENV CGO_ENABLED 0

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

WORKDIR /opt/auth

COPY src/go.mod src/go.sum ./
RUN go mod download && go mod verify

COPY src/. .
RUN go build -v -o /app ./cmd/...


FROM debian:bookworm-slim

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

COPY --from=builder /app /app

EXPOSE 8000
CMD ["/app"]
