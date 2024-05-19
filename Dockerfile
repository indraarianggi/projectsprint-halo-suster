FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY . . 

ENV GOARCH=amd64
ENV GOOS=linux 

RUN env GOARCH=amd64 GOOS=linux go build -o main ./cmd/main.go

FROM alpine
WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8080

CMD ["/app/main"]