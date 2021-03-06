package hamming

import "errors"

const testVersion = 6

var err = errors.New("mismatch in length")

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, err
	}
	distance := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			distance++
		}
	}
	return distance, nil
}
