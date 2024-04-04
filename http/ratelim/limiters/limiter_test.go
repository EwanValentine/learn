package limiters

import (
	"ratelim/internal/redis"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFixedWindowCounter(t *testing.T) {
	adapter := NewInMemoryFixedWindowCounter(1)
	rl := NewLimiter(adapter)
	require.True(t, rl.IsAllowed("client1"))
	require.True(t, rl.IsAllowed("client1"))
	require.False(t, rl.IsAllowed("client1"))
}

func TestFixedWindowCounterRedis(t *testing.T) {
	client, err := redis.New()
	require.NoError(t, err)

	adapter := NewRedisFixedWindowCounterAdapter(client, 1)
	rl := NewLimiter(adapter)

	require.True(t, rl.IsAllowed("client1"))
	require.False(t, rl.IsAllowed("client1"))
}
