FROM golang:1.17.0-alpine AS go-builder
ENV GO111MODULE=on
WORKDIR /go/src/app

RUN sed -e 's/dl-cdn[.]alpinelinux.org/mirrors.tuna.tsinghua.edu.cn/g' -i~ /etc/apk/repositories

# add git for go modules
RUN apk update && apk add make git
RUN go env -w GOPROXY=https://goproxy.io
COPY . .
RUN make manager

FROM alpine
LABEL MAINTAINER="yushk <945009587@qq.com>"

COPY --from=go-builder /go/src/app/manager/data/ /sport/certs/
COPY --from=go-builder /go/src/app/manager/bin/ /sport/

CMD ["/sport/manager"]
