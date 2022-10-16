FROM golang:1.17.13-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

RUN go mod vendor

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]
