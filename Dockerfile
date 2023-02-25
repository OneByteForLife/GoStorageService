FROM golang:1.19 AS builder

WORKDIR /service

COPY . .

RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -ldflags '-extldflags "-static"' -o /go-storage ./cmd/app/main.go


FROM alpine:3.10
COPY --from=builder /go-storage /bin

# Так как при конфигурации сервера я использовал параметр 
#   Prefork чистать тут -> https://docs.gofiber.io/api/fiber, то требуется использовать CMD.
CMD ["sh", "-c", "/bin/go-storage"]