package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestFixedWindowCounter(t *testing.T) {
	rl := NewInMemoryFixedWindowCounter(1)
	require.True(t, rl.IsAllowed("client1"))
	require.True(t, rl.IsAllowed("client1"))
	require.False(t, rl.IsAllowed("client1"))
}
