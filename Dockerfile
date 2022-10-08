FROM golang:1.17-alpine as builder

WORKDIR /app

ENV GO111MODULE=on \
    GOPROXY=https://goproxy.cn,direct \
    CGO_ENABLED=0

# cache
COPY go.mod go.mod
COPY go.sum go.sum
RUN go mod download

COPY . .
RUN go build -trimpath -ldflags '-w -s' -o pure-live .

RUN sed -i 's/dl-cdn.alpinelinux.org/mirrors.aliyun.com/g' /etc/apk/repositories && \
  apk add --no-cache ca-certificates

RUN mkdir build && cp pure-live build && cp -r static build \
    && mkdir build/config && mv config/server.yaml.example build/config/server.yaml \
    && mv config/account.yaml.example build/config/account.yaml

FROM scratch

COPY --from=builder /app/build /
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/

EXPOSE 8800

ENTRYPOINT ["/pure-live","run"]
