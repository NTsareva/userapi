FROM golang:1.22-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o userapi cmd/userservice/main.go

EXPOSE 8081

CMD ["./userapi"]
