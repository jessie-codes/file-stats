FROM golang:latest as builder
COPY . /go/src/github.com/jessie-codes/file-stats
WORKDIR /go/src/github.com/jessie-codes/file-stats
RUN go get ./...
RUN go build -o main .
CMD ["./main", "-i",  "/data/files/*", "-k" ,"/data/keywords", "-o", "/data/stats"]
