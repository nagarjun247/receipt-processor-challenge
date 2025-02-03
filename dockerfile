# Use official Go image
FROM golang:1.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o receipt-processor main.go

# Use a minimal base image
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/receipt-processor .

EXPOSE 8080
CMD ["./receipt-processor"]
