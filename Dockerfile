FROM golang:1.20-alpine3.17 AS build

RUN apk add -U --no-cache git upx curl
ADD . /go/src/app
WORKDIR /go/src/app

# Flags -s -w are to remove debug information from binary.
# -extldflags '-static' is used to statically link binary even with CGO enabled.
RUN GOOS=linux go build -ldflags "-s -w -extldflags '-static'" -o /go/bin/main ./cmd
RUN upx /go/bin/main
RUN apk del git upx

# TODO: improve image using "scratch" and multi stage build image to reduce size

USER 10001
ENTRYPOINT ["/go/bin/main"]
EXPOSE 8090
