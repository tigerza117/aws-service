# Bild stage
FROM golang:1.19.4-alpine3.17 AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/app/*.go

# Run stage
FROM alpine:3.17
WORKDIR /app
COPY --from=builder /app/cmd/app/.env .
COPY --from=builder /app/main .

EXPOSE 3003
CMD [ "/app/main" ]