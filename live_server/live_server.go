package live_server

import (
	"encoding/json"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/websocket"
)

var node_connect, db_server_connect bool
var node_disconnect = make(chan bool)
var db_server_disconnect = make(chan bool)

func (server *LiveServer) connectNode() {
	u_node := url.URL{Scheme: "wss", Host: server.Addr_connected_node, Path: "/websocket"}
	c_node, _, err := websocket.DefaultDialer.Dial(u_node.String(), nil)
	if err != nil {
		log.Print("Cannot connect to Aura node:", err)
		node_connect = false
		return
	}
	err = c_node.WriteMessage(websocket.TextMessage, []byte("{\"jsonrpc\": \"2.0\",\"method\":\"subscribe\",\"id\": 0,\"params\": {\"query\": \"tm.event='Tx' AND unbond.amount>200\"}}"))
	if err != nil {
		log.Println("Error while send data to Aura node, write close:", err)
		node_connect = false
		return
	}
	node_connect = true
	server.node_connection = c_node
	log.Println("Connected to Node")
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
	log.Println("Connected to Forwarder")
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
	log.Println("Stopped connection to Forwarder")
}

func (server *LiveServer) Run() {
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	go func() {
		// connect to node
		for {
			log.Println("Trying to connect to node ...")
			server.connectNode()
			if status := node_connect; status == true {
				break
			}
			time.Sleep(2 * time.Second)
		}

		// connect to Forwarder
		for {
			log.Println("Trying to connect to forwarder ...")
			server.connectDBServer()
			if status := db_server_connect; status == true {
				break
			}
			time.Sleep(2 * time.Second)
		}

		// start websockets
		go server.startWebsocketDBServer()
		go server.startWebsocketNode()
	}()
	defer func() {
		if server.dbserver_connection != nil {
			server.dbserver_connection.Close()
		}
	}()
	defer func() {
		if server.node_connection != nil {
			server.node_connection.Close()
		}
	}()
	// listen event
	for {
		select {
		case <-db_server_disconnect:
			// when disconnect to forwarder, stop connect to node
			server.stopWebsocketNode()
			// restart all
			server.Run()

		case <-node_disconnect:
			// when disconnect to node, stop connect to Forwarder
			server.stopWebsocketDBServer()
			// try to reconnect node
			server.Run()
		case message := <-server.Data_channel:
			log.Println(message)
			var messageForm Message
			err := json.Unmarshal([]byte(message), &messageForm)
			if err != nil {
				log.Fatal(err)
			}
			if messageForm.Result.Data.Value.TxResult.Result.Log != "" {
				var (
					logs []Log
					data = make(map[string]string)
				)
				json.Unmarshal([]byte(messageForm.Result.Data.Value.TxResult.Result.Log), &logs)
				for _, events := range logs {
					for _, event := range events.Events {
						if event.Type == "unbond" {
							for _, attr := range event.Attributes {
								data[attr.Key] = attr.Value
							}
							data["time"] = time.Now().Format("01-02-2006 15:04:05")
							data["delegator"]=messageForm.Result.Events["transfer.sender"][0]
							data["tx_hash"]=messageForm.Result.Events["tx.hash"][0]
							data["tx_fee"]=messageForm.Result.Events["tx.fee"][0]
						}
					}
				}
				jsonData, err := json.Marshal(data)
				if err != nil {
					log.Println("Error when convert data")
				}
				err = server.dbserver_connection.WriteMessage(websocket.TextMessage, []byte(string(jsonData)))
				if err != nil {
					log.Println("Database Server maybe seem down:", err)
				}
			}
		case <-interrupt:
			log.Println("interrupt")

			// Cleanly close the connection by sending a close message and then
			// waiting (with timeout) for the server to close the connection.
			if server.dbserver_connection != nil {
				server.stopWebsocketDBServer()
			}
			if server.node_connection != nil {
				server.stopWebsocketNode()
			}
			return
		}
	}
}
