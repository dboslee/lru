package internal_test

import (
	"testing"

	"github.com/dboslee/lru/internal"

	"github.com/stretchr/testify/require"
)

func TestPushRemove(t *testing.T) {
	ll := internal.NewList[int]()
	length := 10

	for i := 1; i <= length; i++ {
		ll.PushFront(i)
		require.Equal(t, ll.Len(), i)
	}

	for i := length; i <= 1; i++ {
		ll.Remove(ll.Back())
		require.Equal(t, ll.Len(), i)
	}
}

func TestMoveToFront(t *testing.T) {
	ll := internal.NewList[int]()
	e := ll.PushFront(0)
	ll.PushFront(1)
	require.Equal(t, e, ll.Back())

	ll.MoveToFront(e)
	require.NotEqual(t, e, ll.Back())
}

func TestInit(t *testing.T) {
	ll := internal.NewList[int]()

	ll.PushFront(1)
	require.Equal(t, ll.Len(), 1)

	ll.Init()
	require.Equal(t, ll.Len(), 0)
}
