FROM golang:1.23-alpine3.20 AS builder

WORKDIR /app

COPY go.mod go.sum /app/

RUN go mod download

COPY main.go .

RUN go build -o soma .

FROM alpine

WORKDIR /app

COPY --from=builder /app/soma /app/soma

ENTRYPOINT ["./soma"]
CMD ["1 2"]