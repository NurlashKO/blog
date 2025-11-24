FROM golang:1.23.4-bookworm AS builder

ENV CGO_ENABLED=0
ENV GOTOOLCHAIN=auto

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

WORKDIR /opt/blog

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go install github.com/a-h/templ/cmd/templ@latest

COPY . .
RUN templ generate
RUN go build -v -o /app ./cmd/...


FROM debian:bookworm-slim

RUN apt update && \
    apt install -y --no-install-recommends ca-certificates && \
    rm -rf /var/lib/apt/lists/* && \
    update-ca-certificates

COPY --from=builder /app /app

EXPOSE 8000
CMD ["/app"]