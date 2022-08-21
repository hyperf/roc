FROM golang:1.18 as builder

LABEL maintainer="limx <l@hyperf.io>"

WORKDIR /go/cache

ADD go.mod .
ADD go.sum .
RUN go mod download

WORKDIR /go/builder

ADD . .

RUN CGO_ENABLED=0 go build -ldflags="-s -w" -installsuffix cgo -o main ./examples/json_rpc/main.go

FROM scratch

ENV GIN_MODE=release

COPY --from=builder /usr/share/zoneinfo/Asia/Shanghai /etc/localtime
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/builder /

EXPOSE 9501

ENTRYPOINT ["/main"]