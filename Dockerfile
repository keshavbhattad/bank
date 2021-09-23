#Stage 1
FROM golang:1.16-alpine3.13 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go

#Stage 2
FROM alpine:3.13
WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .

EXPOSE 9000
CMD ["/app/main"]