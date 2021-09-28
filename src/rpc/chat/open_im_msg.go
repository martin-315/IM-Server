package main

import (
	rpcChat "Open_IM/src/rpc/chat/chat"
	"flag"
)

func main() {
	rpcPort := flag.Int("port", 300, "rpc listening port")
	flag.Parse()
	rpcServer := rpcChat.NewRpcChatServer(*rpcPort)
	rpcServer.Run()
}
