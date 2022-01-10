# LRU Cache
A simple LRU cache using go generics.

## Examples
Basic usage.
```go
func main() {
    cache := lru.New[string, string]()
    cache.Set("key", "value")
    value := cache.Get("key")
    fmt.Println(value)
}
```
Set the capacity using the `lru.WithCapacity` option. The default capacity is set to 10000.
```go
func main() {
    cache := lru.New[string, string](lru.WithCapacity(100))
    ...
}
```
A thread safe implementation is included for convenience.
```go
func main() {
    cache := lru.NewSync[string, string](lru.WithCapacity(100))
    ...
}
```