# build everything
FROM golang:alpine as builder

ARG source_file

COPY *.go /go/src/
WORKDIR /go/bin
RUN go build -o /go/bin/out /go/src/$source_file

FROM alpine
COPY --from=builder /go/bin/out /usr/local/bin/out

CMD ["/usr/local/bin/out"]
