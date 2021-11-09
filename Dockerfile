FROM golang:alpine

WORKDIR /app

# Download Go modules
COPY go.mod .
RUN go mod download

COPY *.go ./

RUN ls -la
RUN go build -o /netaddr

EXPOSE 8080

CMD ["/netaddr"]
