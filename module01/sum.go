package module01

// Sum will sum up all of the numbers passed
// in and return the result
func Sum(numbers []int) int {
	var ret int
	for _, num := range numbers {
		ret += num
	}
	return ret
}
