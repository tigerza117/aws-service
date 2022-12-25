# Bild stage
FROM golang:1.19.4-alpine3.17 AS builder
WORKDIR /transaction-consumer-lambda
COPY . .
RUN go build -o main cmd/transaction-consumer-lambda/*.go

# Run stage
FROM alpine:3.17
WORKDIR /transaction-consumer-lambda
COPY --from=builder /transaction-consumer-lambda/cmd/transaction-consumer-lambda/.env .
COPY --from=builder /transaction-consumer-lambda/main .

EXPOSE 3003
CMD [ "/transaction-consumer-lambda/main" ]