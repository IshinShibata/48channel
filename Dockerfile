FROM golang:latest

COPY go.mod /server/go.mod
COPY go.sum /server/go.sum
WORKDIR /server
RUN ls
RUN go mod download

COPY . /server

CMD go run main.go
