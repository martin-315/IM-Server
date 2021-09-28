package main

import (
	rpcAuth "Open_IM/src/rpc/auth/auth"
	"flag"
)

func main() {
	rpcPort := flag.Int("port", 10500, "RpcToken default listen port 10800")
	flag.Parse()
	rpcServer := rpcAuth.NewRpcAuthServer(*rpcPort)
	rpcServer.Run()
}
