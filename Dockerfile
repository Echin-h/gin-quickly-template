FROM golang:1.22 AS builder

COPY . /build

WORKDIR /build

RUN set -ex \
    && GO111MODULE=auto CGO_ENABLED=0 go build -ldflags "-s -w -extldflags '-static' -X 'gin-quickly-template/pkg/ \
    version.SysVersion=$(git show -s --format=%h)'" -o App

FROM alpine:latest

WORKDIR /Serve

COPY --from=builder /build/App ./App
# plz replace with true host
COPY --from=builder /build/config.yaml ./config.yaml


RUN apk update && apk add tzdata \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone


ENTRYPOINT ["/Serve/App","server"]

