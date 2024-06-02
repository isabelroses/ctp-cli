package query

import "strings"

func mapValues[K comparable, V any](m map[K]V) []V {
	out := make([]V, 0, len(m))
	for _, tx := range m {
		out = append(out, tx)
	}

	return out
}

func executeFilter[T any](input []T, cb func(T) bool) []T {
	out := make([]T, 0)
	for _, value := range input {
		if cb(value) {
			out = append(out, value)
		}
	}

	return out
}

func SortPortLexicographically(a, b Port) int {
	return strings.Compare(a.Name, b.Name)
}
