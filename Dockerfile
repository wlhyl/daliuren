FROM golang:1.18.3-alpine AS build

# RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories

# RUN apk add --no-cache git gcc make libc-dev
RUN apk update && apk add git gcc make libc-dev

RUN mkdir -p /tmp/swe && \
cd /tmp/swe && \
wget https://www.astro.com/ftp/swisseph/swe_unix_src_2.10.02.tar.gz && \
  tar xvzf swe_unix_src_2.10.02.tar.gz && \
  cd src && \
  make libswe.a

WORKDIR /workspace

# ENV GOPROXY="https://goproxy.cn,direct"

COPY . .

RUN go mod download

RUN go install github.com/swaggo/swag/cmd/swag@latest

RUN swag init

# alpine的git与go 1.18.3不兼容，如果源代码目录中有.git，会引起报错，删除.git目录
RUN rm -rf .git
RUN CGO_LDFLAGS="-L/tmp/swe/src -lswe -lm -ldl -static" go build -o daliuren_server -ldflags '-s -w -extldflags "-static"' .

FROM alpine:3.16.0

# RUN apk add --no-cache ca-certificates

COPY --from=build /workspace/daliuren_server /usr/local/bin/daliuren_server

ENTRYPOINT ["daliuren_server"]
