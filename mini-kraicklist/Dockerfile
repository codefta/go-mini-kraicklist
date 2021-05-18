FROM golang:1.16.4-alpine

RUN apk update && apk add --no-cache git

WORKDIR /app

COPY . .

RUN go mod tidy

RUN go build -o main ./cmd/rest

ENTRYPOINT ["/app/main"]