FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/Bikeleasing-Service/kafka-consumer-test-go

COPY . .

RUN cd src/ && \
    go clean && \
    go install && \
    go mod tidy && \
    go mod verify && \
    mkdir ../dist && \
    CGO_ENABLED=0 go build -a -o ../dist/consumer-test . && \
    cd ../dist \

CMD ["./consumer-test"]
