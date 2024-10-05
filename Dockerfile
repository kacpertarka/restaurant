FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* .

RUN go mod download

COPY . .

RUN mkdir -p /app/bin

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

