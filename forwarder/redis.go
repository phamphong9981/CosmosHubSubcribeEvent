package forwarder

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func saveToRedis(json_data string) {
	var mapData map[string]string
	if err := json.Unmarshal([]byte(json_data), &mapData); err != nil {
		fmt.Println(err)
	}
	log.Println("amount:", mapData["amount"], ",validator address:", mapData["validator"])
	rdb.LPush(ctx, "all", fmt.Sprint(`{"validator": "`, mapData["validator"], `", "time": "`, mapData["time"], `", "amount": "`, mapData["amount"],`", "delegator": "`, mapData["delegator"],`", "tx_hash": "`, mapData["tx_hash"], `"}`))
	rdb.LPush(ctx, mapData["delegator"], fmt.Sprint(`{"time": "`, mapData["time"], `", "amount": "`, mapData["amount"],`", "validator": "`, mapData["validator"],`", "tx_hash": "`, mapData["tx_hash"], `"}`))
}

func publishToRedis(json_data string) {
	var mapData map[string]string
	if err := json.Unmarshal([]byte(json_data), &mapData); err != nil {
		fmt.Println(err)
	}
	err1 := rdb.Publish(ctx, mapData["delegator"], fmt.Sprint(`{"validator": "`, mapData["validator"], `", "time": "`, mapData["time"], `", "amount": "`, mapData["amount"],`", "delegator": "`, mapData["delegator"],`", "tx_hash": "`, mapData["tx_hash"], `"}`)).Err()
	err2 := rdb.Publish(ctx, "all", fmt.Sprint(`{"validator": "`, mapData["validator"], `", "time": "`, mapData["time"], `", "amount": "`, mapData["amount"],`", "delegator": "`, mapData["delegator"],`", "tx_hash": "`, mapData["tx_hash"], `"}`)).Err()
	if err1 != nil || err2 != nil {
		panic(err1)
	}
}
