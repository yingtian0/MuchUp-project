package authz

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/redis/go-redis/v9"
)

var unlockScript = redis.NewScript(`
if redis.call("GET", KEYS[1]) == ARGV[1] then
    return redis.call("DEL", KEYS[1])
else
    return 0
end
`)

func LockProcess(ctx context.Context, client *redis.Client, key string, ttl time.Duration) (string, bool, error) {

	value := uuid.NewString()

	ok, err := client.SetNX(ctx, key, value, ttl).Result()
	if err != nil {
		return "", false, fmt.Errorf("failed to lock redis process: %w", err)
	}
	if !ok {
		return "", false, nil
	}

	return value, true, nil
}

// UnlockProcess releases a Redis-backed lock when the value matches.
func UnlockProcess(ctx context.Context, client *redis.Client, key, value string) error {
	_, err := unlockScript.Run(ctx, client, []string{key}, value).Result()
	if errors.Is(err, redis.Nil) {
		return nil
	}

	return err
}
