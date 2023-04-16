// Package util contains generic things that really should be implemented in
//	core go.
package util

func Map[T, V any](tx func(T) V, input []T) []V {
	output := make([]V, len(input))
	for idx, el := range input {
		output[idx] = tx(el)
	}
	return output
}