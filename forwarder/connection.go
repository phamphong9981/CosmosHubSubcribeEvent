package forwarder

import (
	"sync"

	"github.com/gorilla/websocket"
)

var lock = &sync.Mutex{}

var connectionsList = make(map[*websocket.Conn]string)

func getConnectionsList() map[*websocket.Conn]string {
	lock.Lock()

	defer lock.Unlock()
	return connectionsList
}

func connectClient(connection *websocket.Conn) {
	lock.Lock()
	(connectionsList)[connection] = connection.RemoteAddr().String()
	lock.Unlock()
}

func disconnectClient(connection *websocket.Conn) {
	lock.Lock()
	delete(connectionsList, connection)
	lock.Unlock()
}
