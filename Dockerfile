FROM golang:1.13.8 as builder
WORKDIR /go/src/app
COPY . .
RUN CGO_ENABLED=0 make

FROM alpine:3.11.3
RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.ustc.edu.cn/g' /etc/apk/repositories \
    && apk update \
    && apk add tzdata \
    && cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && mkdir -p /etc/kingshard

COPY --from=builder /go/src/app/bin/kingshard /usr/local/bin/
COPY --from=builder /go/src/app/etc/ks.yaml /etc/kingshard/
CMD kingshard -config=/etc/kingshard/ks.yaml
