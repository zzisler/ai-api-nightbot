FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o ai-api-nightbot ./cmd/ai-api-nightbot/main.go

FROM alpine:latest
WORKDIR /app
COPY --from=builder /app/ai-api-nightbot .
COPY --from=builder /app/internal/prompt ./internal/prompt
CMD ["./ai-api-nightbot"]