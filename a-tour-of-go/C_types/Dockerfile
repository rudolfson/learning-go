# build everything
FROM golang:alpine as builder

COPY *.go /go/src/
WORKDIR /go/bin
RUN for f in $(find /go/src -name \*.go); do GOBIN=/src/bin go build $f; done

# build an image containing the binaries only
FROM alpine

COPY --from=builder /go/bin/ /usr/local/bin/

