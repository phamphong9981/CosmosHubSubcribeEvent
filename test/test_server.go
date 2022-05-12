package test

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)
var client_connect=make(chan bool)
var upgrader = websocket.Upgrader{} // use default options
func echo(w http.ResponseWriter, r *http.Request) {
	
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	client_connect<-true
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			client_connect<-false
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

func Test_server() {
	var addr = flag.String("addr", "localhost:8080", "http service address")
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", echo)
	go func() {
		for{
			select{
			case status:=<-client_connect:
				if !status {
					log.Println("Event node down")
				} else {
					log.Println("Event node up")
				}
			}
		}
	}()
	log.Fatal(http.ListenAndServe(*addr, nil))
}