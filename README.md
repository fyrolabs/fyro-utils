# Fyro-Utils

Set of useful Go functions 

## Ptr
Generic function to convert any value, including function results to pointers.
```go
func Ptr[T any](value T) *T
```

## Slices Map
Similar to array / slice map functions in other languages.

```go
func Map[S any, T any](items []S, cb func(S) T) []T
```


