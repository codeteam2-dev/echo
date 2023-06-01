FROM golang:1.16.3-alpine3.13 AS builder
WORKDIR /go/src/github.com/codeteam2-dev/echo
COPY . .
RUN go build -o /go/bin/echo

FROM scratch
COPY --from=builder /go/bin/echo /go/bin/echo
ENTRYPOINT ["/go/bin/echo"]
