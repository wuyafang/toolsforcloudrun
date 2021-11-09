FROM golang:alpine

RUN ls -la src
RUN go build -o netaddr main.go

CMD ["./netaddr"]
