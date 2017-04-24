package leap

const testVersion = 3

func IsLeapYear(year int) bool {
	if (year % 4) != 0 {
		return false
	}
	mod := year % 100
	if mod != 0 {
		return true
	}
	if (year % 400) == 0 {
		return true
	}
	return false
}
