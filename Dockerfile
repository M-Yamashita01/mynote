FROM golang:1.17.13-alpine3.16

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go install github.com/cosmtrek/air@latest

RUN go mod download

# COPY ./ ./

# RUN go build -o /mynote /app/cmd/mynote/main.go

EXPOSE 8080

# CMD [ "/mynote" ]
CMD ["air", "-c", ".air.toml"]
