FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN mkdir -p logs && touch logs/app.log && chmod 666 logs/app.log

RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go 

FROM alpine:latest

WORKDIR /app

# Myapp, .env, logs — mavjud
COPY --from=builder /app/myapp .
COPY --from=builder /app/.env .
COPY --from=builder /app/logs ./logs

# ✅ Swagger docs papkasini ham nusxalaymiz
COPY --from=builder /app/docs ./docs

EXPOSE 8080

CMD ["./myapp"]
