package limiters

import (
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
