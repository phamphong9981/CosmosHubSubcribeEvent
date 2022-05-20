package main
import (
	"datn/db_server"
)
func main(){
	var db_sv=db_server.DBServer{Addr_connected_live_server: "localhost:8080",Addr_connected_recover_server: ""}
	db_sv.Run()
}