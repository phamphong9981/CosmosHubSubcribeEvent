package db_server

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

type DBServer struct {
	Addr_connected_live_server    string
	Addr_connected_recover_server string
}

func echo(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Print("Node", r.URL, r.URL.User, r.Host, "up")
	defer c.Close()
	connectClient(c)
	log.Print(connectionsList)
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("Node down:", err)
			disconnectClient(c)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}

func (server *DBServer) Run() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", echo)
	log.Fatal(http.ListenAndServe(server.Addr_connected_live_server, nil))
}
