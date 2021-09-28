package main

import (
	"Open_IM/src/rpc/friend/friend"
	"flag"
)

func main() {

	rpcPort := flag.Int("port", 10300, "get RpcFriendPort from cmd,default 12000 as port")
	flag.Parse()
	rpcServer := friend.NewFriendServer(*rpcPort)
	rpcServer.Run()
}
