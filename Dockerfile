# Build from the backend workspace: docker build -f api-gateway/Dockerfile .
FROM golang:1.26.2-alpine AS builder
WORKDIR /build
COPY api-gateway ./api-gateway
COPY services/auth/api ./services/auth/api
COPY common/tokens ./common/tokens
WORKDIR /build/api-gateway
RUN CGO_ENABLED=0 go build -trimpath -ldflags "-s -w" -o /out/gateway ./cmd

FROM alpine:3.23
RUN apk add --no-cache ca-certificates && adduser -D -u 10001 app
COPY --from=builder /out/gateway /bin/gateway
USER app
ENTRYPOINT ["/bin/gateway"]
