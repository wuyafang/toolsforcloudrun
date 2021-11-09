FROM golang:alpine

RUN ls -la
RUN go build -o netaddr main.go

CMD ["./netaddr"]
