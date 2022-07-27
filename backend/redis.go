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

// func getUnbondFromDelegator(delegator string) []string {
// 	if delegator == "" {
// 		return nil
// 	}
// 	if delegator == "all" {
// 		return rdb.LRange(ctx, "all", 0, 20).Val()
// 	}
// 	return rdb.LRange(ctx, delegator, 0, -1).Val()
// }

func subscibeInit() {
	subscribeAllChannel = rdb.PSubscribe(ctx, "*")
}
func handleSubscibe() {

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
