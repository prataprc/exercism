package pangram

const testVersion = 1

func IsPangram(s string) bool {
	var table [256]byte
	for _, x := range s {
		table[byte(x)] = 1
	}
	for i := 65; i < (65 + 26); i++ {
		if (table[i] + table[i+32]) == 0 {
			return false
		}
	}
	return true
}
