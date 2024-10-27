FROM golang:alpine AS builder
WORKDIR /var/www/app
COPY . .
RUN go mod download && go build -o app cmd/telegram_bot/main.go

FROM alpine:latest AS app
WORKDIR /var/www/app
COPY --from=builder /var/www/app/app .
CMD [ "./app" ]