package forwarder

import (
	"flag"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type DBServer struct {
	WebsocketAddress string
}

func echo(w http.ResponseWriter, r *http.Request) {
	var upgrader = websocket.Upgrader{} // use default options
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Print("Node", c.RemoteAddr().String(), " up")
	

	connectClient(c)
	log.Print(connectionsList)
}

func handleConnection() {
	for {
		for c, addr := range getConnectionsList() {
			log.Print("Connected to live server "+ addr)
			defer c.Close()
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					log.Println("Node ", addr, " down:", err)
					disconnectClient(c)
					break
				}
				saveToRedis(string(message))
				publishToRedis(string(message))
				saveToMongo(string(message))
			}
		}
		// There arent any available live sever which could connect
		log.Print("There arent any available live sever which could connect. Wait 1 second for live server up")
		time.Sleep(1*time.Second)
	}
}

func (server *DBServer) Run() {
	flag.Parse()
	log.SetFlags(0)
	go handleConnection()
	http.HandleFunc("/websocket", echo)
	log.Fatal(http.ListenAndServe(server.WebsocketAddress, nil))
}
