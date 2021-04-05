FROM golang:1.16

WORKDIR home

COPY . .

RUN go get ./...

RUN go build

CMD ["./chaos_monkey"]