FROM golang:1.22.1-bookworm as builder

ENV CGO_ENABLED 0

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

COPY --from=builder /app /app

CMD ["/app"]