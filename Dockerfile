# stage 1
FROM golang:1.26-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main ./cmd/api

# stage 2
FROM alpine:3.20

RUN adduser -D -g '' appuser

WORKDIR /app

COPY --from=builder /app/main .

USER appuser

EXPOSE 5000

ENTRYPOINT ["./main"]