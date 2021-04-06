FROM golang:1.16 AS builder

WORKDIR /build

COPY . .

RUN go get ./...

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o chaos_monkey .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR home
COPY --from=builder /build/chaos_monkey .

CMD ["./chaos_monkey"]