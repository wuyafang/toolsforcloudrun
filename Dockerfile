FROM golang:alpine

RUN go build -o netaddr main.go

CMD ["./netaddr"]
