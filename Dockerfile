FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/Bikeleasing-Service/kafka-consumer-test-go
COPY . .
WORKDIR ./src
RUN go clean && go install
RUN go mod tidy
RUN go mod verify
RUN mkdir ../dist
RUN CGO_ENABLED=0 go build -a -o ../dist/consumer-test .
WORKDIR ../dist
CMD ["./consumer-test"]
