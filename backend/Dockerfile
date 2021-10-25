# run `docker build -t backend .`
# set the base image
FROM golang:1.15-alpine3.14
RUN apk add bash ca-certificates git gcc g++ libc-dev
RUN mkdir /app
ADD . /app/
WORKDIR /app
ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct
RUN go mod download
RUN go build -v
RUN ls