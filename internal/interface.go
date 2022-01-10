package internal

type LRU[K comparable, V any] interface {
	Get(K) (V, bool)
	Set(K, V)
	Delete(K) bool
	Peek(K) (V, bool)
	Flush()
	Len() int
}
