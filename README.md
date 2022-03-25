# Channel driven multiplexing connection

[![Test](https://github.com/hyperf/roc/actions/workflows/test.yml/badge.svg)](https://github.com/hyperf/roc/actions/workflows/test.yml)

## How to install

```shell
go get github.com/hyperf/roc
```

## How to use

```go
package main

import (
	"fmt"
	"github.com/hyperf/roc"
	"github.com/hyperf/roc/server"
	"net"
)

func main() {
	serv := server.NewTcpServer("127.0.0.1:9501", func(conn net.Conn, packet *roc.Packet, server *server.TcpServer) {
		id := packet.GetId()
		body := packet.GetBody()

		fmt.Println(id, body)
	})

	serv.Start()
}
```