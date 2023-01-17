FROM golang:1.19-alpine AS builder

WORKDIR /go/src/github.com/Bikeleasing-Service/kafka-consumer-test-go

COPY . .

ENV SRC_DIR "./src/"
ENV DIST_DIR "./dist"

WORKDIR "${SRC_DIR}"

RUN go clean && \
    go install && \
    go mod tidy && \
    go mod verify && \
    mkdir "${DIST_DIR}" && \
    CGO_ENABLED=0 go build -a -o "${DIST_DIR}/consumer-test" .

WORKDIR "${DIST_DIR}/"

CMD ["./consumer-test"]
