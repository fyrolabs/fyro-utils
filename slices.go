package util

func Map[S any, T any](items []S, cb func(S) T) []T {
	results := make([]T, len(items))

	for i, item := range items {
		results[i] = cb(item)
	}
	return results
}
