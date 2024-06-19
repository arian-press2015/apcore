# Stage 1: Build the application
FROM golang:1.22.2-alpine AS builder

RUN apk update && apk add --no-cache \
    git \
    gcc \
    musl-dev

WORKDIR /app

COPY go.mod go.sum ./

ENV GOPROXY=http://localhost:8081,direct
RUN go mod download

COPY . .

RUN go install github.com/swaggo/swag/cmd/swag@latest
RUN swag init --parseDependency

RUN go build -o main .

# Stage 2: Create a minimal runtime image
FROM alpine:3.20.0

WORKDIR /app

COPY --from=builder /app/main .

# COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./main"]
