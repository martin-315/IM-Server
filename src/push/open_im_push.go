package main

import (
	"Open_IM/src/push/logic"
	"flag"
	"sync"
)

func main() {
	rpcPort := flag.Int("port", 400, "rpc listening port")
	flag.Parse()
	var wg sync.WaitGroup
	wg.Add(1)
	logic.Init(*rpcPort)
	logic.Run()
	wg.Wait()
}
