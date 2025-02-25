FROM golang:alpine

WORKDIR /jwt-auth

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o myapp cmd/main.go

EXPOSE 8080

CMD ["./myapp"]