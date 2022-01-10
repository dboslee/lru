package lru_test

import (
	"testing"

	"github.com/dboslee/lru"
	"github.com/stretchr/testify/require"
)

func TestCapacity(t *testing.T) {
	tests := []struct {
		capacity int
	}{
		{1},
		{10},
		{100},
	}

	lru.New[int, int](lru.WithCapacity(10))
	for _, tc := range tests {
		lru := lru.New[int, int](lru.WithCapacity(tc.capacity))
		for i := 0; i < tc.capacity+1; i++ {
			lru.Set(i, i)
		}

		require.Equal(t, tc.capacity, lru.Len(), "expected capacity to be full")

		_, ok := lru.Get(0)
		require.False(t, ok, "expected key to be evicted")

		_, ok = lru.Get(1)
		require.True(t, ok, "expected key to exist")
	}
}

func TestGetMissing(t *testing.T) {
	lru := lru.New[int, int]()
	_, ok := lru.Get(0)
	require.False(t, ok, "expected not ok")
}

func TestSetGet(t *testing.T) {
	lru := lru.New[int, int]()
	value := 100

	lru.Set(1, value)
	value, ok := lru.Get(1)

	require.True(t, ok, "expected ok")
	require.Equal(t, value, value, "expected set value %s", value)
}

func TestDelete(t *testing.T) {
	lru := lru.New[int, int]()

	key, value := 1, 100
	lru.Set(key, value)
	require.Equal(t, lru.Len(), 1)

	ok := lru.Delete(key)
	require.True(t, ok, "expected ok")
}

func TestDeleteMissing(t *testing.T) {
	lru := lru.New[int, int]()
	key := 100
	ok := lru.Delete(key)
	require.False(t, ok, "expected not ok")
}

func TestFlush(t *testing.T) {
	lru := lru.New[int, int]()
	key, value := 1, 100
	lru.Set(key, value)
	require.Equal(t, lru.Len(), 1)

	lru.Flush()
	require.Equal(t, lru.Len(), 0)

	_, ok := lru.Get(key)
	require.False(t, ok, "expected not ok")
}
