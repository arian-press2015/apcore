FROM golang:1.22.2-alpine

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

EXPOSE 8080

CMD ["/app/main"]
