package main
import (
	"datn/db_server"
)
func main(){
	var db_sv=db_server.DBServer{WebsocketAddress: "localhost:8080"}
	db_sv.Run()
}