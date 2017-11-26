FROM golang:alpine
MAINTAINER Chs97 <chsqaq@gmail.com>
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata
RUN apk update && \
    apk add --no-cache git
WORKDIR /go/src/app
COPY /srv/auth /go/src/app
RUN go get
RUN go build main.go