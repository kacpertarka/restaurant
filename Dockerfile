FROM golang:1.23.1-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.* .

RUN go mod download

RUN mkdir -p /app/bin

COPY . .

EXPOSE 8080

CMD ["air", "-c", ".air.toml"]

