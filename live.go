package main

import (
	"datn/live_server"
)

func main() {
	var live_sv = (live_server.LiveServer{Addr_connected_database: "localhost:8080",Addr_connected_node: "rpc.cosmos.network:443", Data_channel: make(chan string)})
	live_sv.Run()
}
