FROM golang:1.19-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
COPY main.go ./

RUN go mod download

COPY src/ ./src

RUN go build

EXPOSE 8080

CMD ["./go-server"]