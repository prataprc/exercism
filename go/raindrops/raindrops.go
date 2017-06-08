package raindrops

import "strconv"

const testVersion = 3

var atoms = [16]string{
	"", "", "",
	"Pling", "", "Plang",
	"", "Plong", "PlingPlang",
	"", "PlingPlong", "",
	"PlangPlong", "", "",
	"PlingPlangPlong",
}
var divs = [3]int{3, 5, 7}

func Convert(n int) string {
	index := 0
	for _, d := range divs {
		if (n % d) == 0 {
			index += d
		}
	}
	if index == 0 {
		return strconv.Itoa(n)
	}
	return atoms[index]
}
