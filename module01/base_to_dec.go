package module01

import (
	"math"
	"strings"
)

// BaseToDec takes in a number and the base it is currently
// in and returns the decimal equivalent as an integer.
//
// Eg:
//
//   BaseToDec("E", 16) => 14
//   BaseToDec("1110", 2) => 14
//
func BaseToDec(value string, base int) int {
	var ret int
	var chars string = "0123456789ABCDEF"
	for i, v := range strings.Split(Reverse(value), "") {
		ret += strings.Index(chars, v) * int(math.Pow(float64(base), float64(i)))
	}
	return ret
}
