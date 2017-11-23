FROM golang:alpine
RUN apk add --no-cache tzdata && \
    cp /usr/share/zoneinfo/Asia/Shanghai /etc/localtime && echo "Asia/Shanghai" > /etc/timezone && \
    apk del tzdata
RUN apk update && \
    apk add --no-cache git
COPY package.json /tmp/package.json
WORKDIR /go/src/app
COPY . /opt/workdir
RUN go build main.go