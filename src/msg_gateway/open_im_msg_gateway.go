package main

import (
	"Open_IM/src/msg_gateway/gate"
	"flag"
	"sync"
)

func main() {
	rpcPort := flag.Int("rpc_port", 10600, "rpc listening port")
	wsPort := flag.Int("ws_port", 10700, "ws listening port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	gate.Init(*rpcPort, *wsPort)
	gate.Run()
	wg.Wait()
}
