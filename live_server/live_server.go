package live_server

import (
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

type LiveServer struct {
	Addr_connected_database string
	Addr_connected_node     string
	Data_channel            chan string
	node_connection         *websocket.Conn
	dbserver_connection     *websocket.Conn
}

var node_connect, db_server_connect bool
var node_disconnect, db_server_disconnect chan bool

func (server *LiveServer) connectNode() {
	u_node := url.URL{Scheme: "wss", Host: server.Addr_connected_node, Path: "/websocket"}
	c_node, _, err := websocket.DefaultDialer.Dial(u_node.String(), nil)
	if err != nil {
		log.Print("Cannot connect to Aura node:", err)
		node_connect = false
		return
	}
	err = c_node.WriteMessage(websocket.TextMessage, []byte("{\"jsonrpc\": \"2.0\",\"method\":\"subscribe\",\"id\": 0,\"params\": {\"query\": \"tm.event='Tx' AND unbond.validator='cosmosvaloper178h4s6at5v9cd8m9n7ew3hg7k9eh0s6wptxpcn'\"}}"))
	if err != nil {
		log.Println("Error while send data to Aura node, write close:", err)
		node_connect = false
		return
	}
	node_connect = true
	server.node_connection = c_node
}

func (server *LiveServer) connectDBServer() {
	u := url.URL{Scheme: "ws", Host: server.Addr_connected_database, Path: "/websocket"}

	c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Println("Cannot connect to Database Server:", err)
		db_server_connect = false
		return
	}
	db_server_connect = true
	server.dbserver_connection = c
}

func (server *LiveServer) startWebsocketDBServer() {
	for {
		_, message, err := server.dbserver_connection.ReadMessage()
		if err != nil {
			log.Println("Disconnect to Database Server", err)
			db_server_disconnect <- true
			return
		}
		log.Printf("recv: %s", message)
	}
}

func (server *LiveServer) startWebsocketNode() {
	for {
		_, message, err := server.node_connection.ReadMessage()
		if err != nil {
			log.Println("Disconnect to Aura node:", err)
			node_disconnect <- true
			break
		}
		server.Data_channel <- string(message)
	}
}

func (server *LiveServer) stopWebsocketNode() {
	err := server.node_connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("Stop node fail:", err)
		return
	}
	select {
	case <-node_disconnect:
	case <-time.After(time.Second):
	}
	log.Println("Stopped connection to Node")
}

func (server *LiveServer) stopWebsocketDBServer() {
	err := server.dbserver_connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		log.Println("write close:", err)
		return
	}
	select {
	case <-db_server_disconnect:
	case <-time.After(time.Second):
	}
	log.Println("Stopped connection to DB Server")
}

func (server *LiveServer) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connect to node
	for {
		log.Println("Trying to connect to node ...")
		server.connectNode()
		if status := node_connect; status == true {
			break
		}
		time.Sleep(2 * time.Second)
	}

	// connect to db server
	for {
		log.Println("Trying to connect to db server ...")
		server.connectDBServer()
		if status := db_server_connect; status == true {
			break
		}
		time.Sleep(2 * time.Second)
	}

	// start websockets
	go server.startWebsocketDBServer()
	go server.startWebsocketNode()
	defer server.dbserver_connection.Close()
	defer server.node_connection.Close()
	// listen event
	for {
		select {
		case <-db_server_disconnect:
			// when disconnect to db server, stop connect to node
			server.stopWebsocketNode()
			// restart all
			server.Run()

		case <-node_disconnect:
			// when disconnect to node, stop connect to db server
			server.stopWebsocketDBServer()
			// try to reconnect node
			server.Run()
		case message := <-server.Data_channel:
			log.Println(server)
			err := server.dbserver_connection.WriteMessage(websocket.TextMessage, []byte(message))
			if err != nil {
				log.Println("Database Server maybe seem down:", err)
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			server.stopWebsocketNode()
			server.stopWebsocketDBServer()
			return
		}
	}
}
