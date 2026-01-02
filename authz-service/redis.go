package authzservice

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"

	"github.com/redis/go-redis/v9"
)

var (
	ctx = context.Background()
	rdb *redis.Client
)

var unlockScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
`)

func InitRedis () {
	rdb = redis.NewClient(&redis.Options{
		Addr: os.Getenv("REDIS_PORT"),
		Password: "",
		DB: 0,})
}


func LockProccss(rdb redis.Client,key string,ttl time.Duration) (string,bool,error) {

	value := uuid.NewString()
	ok,err := rdb.SetNX(ctx,value,key,ttl).Result()
	if err != nil || !ok {
		return "",false,fmt.Errorf("faild to lock Redis process:%w",err)
	}

	return value,true,nil
}

func UnlockProcess(rdb redis.Client,key,value string) error {
	_, err := unlockScript.Run(ctx, rdb, []string{key}, value).Result()
	return err
}

