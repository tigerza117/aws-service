# Bild stage
FROM golang:1.19.4-alpine3.17 AS builder
WORKDIR /test
COPY . .
RUN go build -o main cmd/test/*.go

# Run stage
FROM alpine:3.17
WORKDIR /test
#COPY --from=builder /test/cmd/transaction-consumer/.env .
COPY --from=builder /test/main .

EXPOSE 3003
CMD [ "/test/main" ]