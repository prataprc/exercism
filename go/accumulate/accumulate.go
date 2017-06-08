package accumulate

const testVersion = 1

func Accumulate(xs []string, fn func(string) string) (acc []string) {
	for _, x := range xs {
		acc = append(acc, fn(x))
	}
	return acc
}
