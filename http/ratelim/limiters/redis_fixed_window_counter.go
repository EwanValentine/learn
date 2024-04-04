package limiters

import (
	"context"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis/v9"
)

type RedisFixedWindowCounterAdapter struct {
	rps    int
	client *redis.Client
}

const luaScript = `
local keyName = KEYS[1]
local rps = tonumber(ARGV[1])
local currentCount = redis.call('GET', keyName)

if currentCount then
	currentCount = redis.call('INCR', keyName)
else
	currentCount = redis.call('SET', keyName, 1, 'EX', rps)
	currentCount = 1
end

return currentCount
`

func NewRedisFixedWindowCounterAdapter(client *redis.Client, rps int) *RedisFixedWindowCounterAdapter {
	return &RedisFixedWindowCounterAdapter{client: client, rps: rps}
}

func (f *RedisFixedWindowCounterAdapter) IsAllowed(id string) bool {
	rps := strconv.Itoa(f.rps)

	currentTime := strconv.Itoa(int(time.Now().Unix()))
	key := id + ":" + currentTime

	res, err := f.client.Eval(context.Background(), luaScript, []string{key}, rps).Result()
	if err != nil {
		log.Println(err)
		return false
	}

	count := res.(int64)

	log.Println(count <= int64(f.rps))

	return count <= int64(f.rps)
}
