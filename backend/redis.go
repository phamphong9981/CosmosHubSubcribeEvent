package backend

import (
	"context"

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

func subscibeAll() {
	subscribeAllChannel = rdb.Subscribe(ctx, "all")
	for {
		msg, err := subscribeAllChannel.ReceiveMessage(ctx)
		list := getConnectionsList()
		if err != nil {
			panic(err)
		}
		for con, _ := range list {
			err = con.WriteMessage(websocket.TextMessage, []byte(msg.Payload))
		}
	}
}
