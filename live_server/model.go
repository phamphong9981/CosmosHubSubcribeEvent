package live_server

import (
	"github.com/gorilla/websocket"
)

type LiveServer struct {
	Addr_connected_database string
	Addr_connected_node     string
	Data_channel            chan string
	node_connection         *websocket.Conn
	dbserver_connection     *websocket.Conn
}

type Message struct {
	Jsonrpc   string
	Id     int
	Result struct {
		Data struct {
			Value struct {
				TxResult struct {
					Result struct {
						Log string
					}
				}
			}
		}
	}
}

type Attribute struct {
	Key   string `json: "key"`
	Value string `json: "value"`
}
type Event struct {
	Type string `json: "type"`
	Attributes []Attribute `json: "attribute"`
}
type Log struct{
	Events []Event `json: "events"`
}
