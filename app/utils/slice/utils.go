package slice

import "math/rand"

func Shuffle[T any](src []T, seed int64) []T {
	r := rand.New(rand.NewSource(seed))
	dest := make([]T, len(src))
	perm := r.Perm(len(src))
	for i, v := range perm {
		dest[v] = src[i]
	}
	return dest
}
