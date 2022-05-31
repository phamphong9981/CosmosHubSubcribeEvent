package backend

import (
	"sync"

	"github.com/gorilla/websocket"
)

var lock = &sync.Mutex{}

var connectionsList = make(map[*websocket.Conn]bool)

func getConnectionsList() map[*websocket.Conn]bool {
	lock.Lock()
	
	defer lock.Unlock()
	return connectionsList
}

func connectClient(connection *websocket.Conn) {
	lock.Lock()
	(connectionsList)[connection] = true
	lock.Unlock()
}

func disconnectClient(connection *websocket.Conn) {
	lock.Lock()
	delete(connectionsList, connection)
	lock.Unlock()
}