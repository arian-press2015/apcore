FROM golang:1.22.2-alpine

WORKDIR /app

COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
ENV GOPROXY="https://goproxy.io,direct"
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["/app/main"]
