package husky

func Filter[V any](xs []V, fn func(V) bool) []V {
	var result []V
	for _, item := range xs {
		if fn(item) {
			result = append(result, item)
		}
	}
	return result
}

func Map[T any, R any](xs []T, fn func(T) R) []R {
	result := make([]R, len(xs))

	for i, x := range xs {
		result[i] = fn(x)
	}

	return result
}

func Foldl[T any, R any](initial R, xs []T, fn func(R, T) R) R {
	for _, x := range xs {
		initial = fn(initial, x)
	}
	return initial
}

func ForEach[T any](xs []T, fn func(T)) {
	for _, x := range xs {
		fn(x)
	}
}

func Elem[T comparable](collection []T, element T) bool {
	for _, item := range collection {
		if item == element {
			return true
		}
	}
	return false
}

func Any[V any](xs []V, fn func(V) bool) bool {
	for _, item := range xs {
		if fn(item) {
			return true
		}
	}
	return false
}

func Take[T any](n int, xs []T) []T {
	return xs[:n]
}

func Not(b bool) bool {
	return !b
}
