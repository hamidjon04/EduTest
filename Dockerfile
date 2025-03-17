FROM golang:1.23.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

# Dastur qismida logs papkasini yaratamiz
RUN mkdir -p logs

# Bu qismda manzilni aniq ko'rsatib, myapp faylini o'rnatamiz
RUN CGO_ENABLED=0 GOOS=linux go build -o myapp ./cmd/main.go 

FROM alpine:latest

WORKDIR /app

# Builder qatlamidan myapp ni nusxalaymiz
COPY --from=builder /app/myapp .

COPY --from=builder /app/.env .

# logs papkasini nusxalaymiz
COPY --from=builder /app/logs ./logs

EXPOSE 8080

CMD ["./myapp"]
