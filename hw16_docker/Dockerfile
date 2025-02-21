FROM golang:1.20-alpine as builder

WORKDIR /app

COPY go.mod .

RUN go mod download

COPY . .

RUN go build -o /app/cmd/server ./server/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/cmd /app

COPY --from=builder /app/.env /app

EXPOSE 8080

CMD ["./server"]