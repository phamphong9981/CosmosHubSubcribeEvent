package db_server

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

}
