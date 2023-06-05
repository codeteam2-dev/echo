FROM golang:1.16.3-alpine3.13 AS builder
WORKDIR /go/src/github.com/codeteam2-dev/echo
COPY . .
RUN go mod download
RUN go build -o /go/bin/echo -tags netgo -installsuffix netgo -ldflags="-w -s"

FROM scratch
LABEL maintainer="codeteam2-dev"
LABEL org.opencontainers.image.source = "https://github.com/codeteam2-dev/echo"
COPY --from=builder /go/bin/echo /echo
ENTRYPOINT ["/echo"]
