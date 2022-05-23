package backend

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func getUnbondFromValidator(validator string) []string {
	if validator=="" {
		return nil
	}
	return rdb.LRange(ctx, validator, 0, -1).Val()
}