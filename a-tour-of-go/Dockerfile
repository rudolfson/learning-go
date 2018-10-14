FROM golang:alpine

COPY *.go /go/src/
WORKDIR bin
# build everything
RUN for f in $(find /go/src -name \*.go); do GOBIN=/src/bin go build $f; done

