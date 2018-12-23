FROM golang:latest
RUN go get github.com/codegangsta/gin
ENV PATH="/go/bin:${PATH}"
WORKDIR /go/src