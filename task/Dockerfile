FROM golang@sha256:9dd2625a1ff2859b8d8b01d8f7822c0f528942fe56cfe7a1e7c38d3b8d72d679 AS builder

WORKDIR /app

COPY . .

RUN go build -o main ./cmd/app/main.go

FROM alpine@sha256:beefdbd8a1da6d2915566fde36db9db0b524eb737fc57cd1367effd16dc0d06d

WORKDIR /app

RUN addgroup -S appgroup && adduser -S appuser -G appgroup

COPY --from=builder /app/main .

RUN chown -R appuser:appgroup /app

USER appuser

CMD ["./main"]
