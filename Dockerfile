FROM golang:1.17.7-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build

CMD ["./app"]

