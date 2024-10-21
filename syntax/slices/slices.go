package slices

func FilterFunc[T any](s []T, f func(e T) bool) []T {
	return filter(s, f)
}

func FilterComparable[T comparable](s []T, cmp T) []T {
	return filter(s, func(e T) bool {
		return e == cmp
	})
}

func filter[T any](s []T, f func(e T) bool) []T {
	var fs []T
	for _, s := range s {
		if f(s) {
			fs = append(fs, s)
		}
	}
	return fs
}

func HasFunc[T any](s []T, f func(e T) bool) bool {
	return has(s, f)
}

func HasComparable[T comparable](s []T, cmp T) bool {
	return has(s, func(e T) bool {
		return e == cmp
	})
}

func has[T any](s []T, f func(e T) bool) bool {
	for _, s := range s {
		if f(s) {
			return true
		}
	}
	return false
}
