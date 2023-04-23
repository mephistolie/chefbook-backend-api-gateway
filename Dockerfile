FROM golang:alpine as builder

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o /main cmd/main.go

FROM alpine:latest

COPY --from=builder main /bin/main
COPY --from=builder build/docs /docs
ENTRYPOINT ["/bin/main"]
