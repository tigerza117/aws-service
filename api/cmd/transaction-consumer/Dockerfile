# Bild stage
FROM golang:1.19.4-alpine3.17 AS builder
WORKDIR /transaction-consumer
COPY . .
RUN go build -o main cmd/transaction-consumer/*.go

# Run stage
FROM alpine:3.17
WORKDIR /transaction-consumer
COPY --from=builder /transaction-consumer/cmd/transaction-consumer/.env .
COPY --from=builder /transaction-consumer/main .

EXPOSE 3003
CMD [ "/transaction-consumer/main" ]