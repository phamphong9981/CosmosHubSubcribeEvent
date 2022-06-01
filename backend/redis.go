package backend

import (
	"context"
	"log"

	"github.com/go-redis/redis/v8"
	"github.com/gorilla/websocket"
)

var ctx = context.Background()
var subscribeAllChannel = new(redis.PubSub)
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getUnbondFromValidator(validator string) []string {
	if validator == "" {
		return nil
	}
	if validator == "all" {
		return rdb.LRange(ctx, "all", 0, 20).Val()
	}
	return rdb.LRange(ctx, validator, 0, -1).Val()
}

func subscibeInit() {
	subscribeAllChannel = rdb.PSubscribe(ctx, "*")
}
func handleSubscibe() {
	// for {
	// 	msg, err := subscribeAllChannel.ReceiveMessage(ctx)
	// 	list := getConnectionsList()
	// 	if err != nil {
	// 		panic(err)
	// 	}
	// 	for con, validator := range list {
	// 		if validator == msg.Channel {
	// 			err = con.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
	// 		}
	// 	}
	// }

	ch := subscribeAllChannel.Channel()

	for msg := range ch {
		list := getConnectionsList()
		for con, validator := range list {
			if validator == msg.Channel {
				err:= con.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
				log.Print(validator)
				if err!=nil {
					log.Print("There are disconnect")
					disconnectClient(con)
				}
			}
		}
	}
}
