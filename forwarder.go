package main
import (
	"datn/forwarder"
)
func main(){
	var db_sv=forwarder.DBServer{WebsocketAddress: "localhost:8080"}
	db_sv.Run()
}