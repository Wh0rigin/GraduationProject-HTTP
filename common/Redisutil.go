package common

import (
	"fmt"

	"github.com/gin-contrib/sessions/redis"
)

func Redisutil() redis.Store {
	store, err := redis.NewStore(10, "tcp", "localhost:6379", "", []byte("secret"))
	if err != nil {
		fmt.Print(err.Error())
	}
	return store
}
