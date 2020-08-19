package module01

import (
	"strconv"
)

// DecToBase will return a string representing
// the provided decimal number in the provided base.
// This is limited to bases 2-16 for simplicity.
//
// Eg:
//
//   DecToBase(14, 16) => "E"
//   DecToBase(14, 2) => "1110"
//
//func DecToBase(dec, base int) string {
//var answer string = ""
//hex := map[int]string{
//10: "A",
//11: "B",
//12: "C",
//13: "D",
//14: "E",
//15: "F",
//}
//var remainder int
//for dec != 0 {
//remainder = dec % base
//if remainder < 10 {
//answer = strconv.Itoa(remainder) + answer
//} else {
//answer = hex[remainder] + answer
//}
//dec = dec / base
//}
//return answer
//}

// Better answer

func DecToBase(dec, base int) string {
	// We know that the question set will not go higher than base 16
	const chars = "0123456789ABCDEF"
	var answer string
	for dec != 0 {
		remainder := dec % base
		// chars can be accessed by index
		answer = string(chars[remainder]) + answer
		dec = dec / base
	}
	return answer
}
