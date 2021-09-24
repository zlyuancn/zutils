package zutils

import (
	"context"

	"github.com/go-redis/redis/v8"
)

var Redis = new(redisUtil)

type redisUtil struct{}

func (*redisUtil) Cas(ctx context.Context, client redis.UniversalClient, key, v1, v2 string) (bool, error) {
	const script = `
  local v = redis.call("get", KEYS[1])

  if (v == KEYS[2]) then
      redis.call("set", KEYS[1], KEYS[3])
      return 1
  end

  return 0`
	return client.Eval(ctx, script, []string{key, v1, v2}).Bool()
}
