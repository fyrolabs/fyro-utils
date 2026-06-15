package utils

func Map[S any, T any](items []S, cb func(S) T) []T {
	results := make([]T, len(items))

	for i, item := range items {
		results[i] = cb(item)
	}
	return results
}

func MapWithErr[S any, T any](items []S, cb func(S) (T, error)) ([]T, error) {
	results := make([]T, len(items))

	for i, item := range items {
		res, err := cb(item)
		if err != nil {
			return nil, err
		}
		results[i] = res
	}
	return results, nil
}
