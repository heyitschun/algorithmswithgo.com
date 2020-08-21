package module01

// Fibonacci returns the nth fibonacci number.
//
// A Fibonacci number N is defined as:
//
//   Fibonacci(N) = Fibonacci(N-1) + Fibonacci(N-2)
//
// Where the following base cases are used:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//
//
// Examples:
//
//   Fibonacci(0) => 0
//   Fibonacci(1) => 1
//   Fibonacci(2) => 1
//   Fibonacci(3) => 2
//   Fibonacci(4) => 3
//   Fibonacci(5) => 5
//   Fibonacci(6) => 8
//   Fibonacci(7) => 13
//   Fibonacci(14) => 377
//
func Fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fibonacci(n-1) + Fibonacci(n-2)
	}
}

// The given solution, while not many lines of code, is inefficient and will
// be very expensive to run for larger numbers because it is calculating
// multiple times. Improvements could be using an iterative approach:

func AltFib(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		nMin1 := 1
		nMin2 := 0
		var currentFib int
		for i := 2; i <= n; i++ {
			currentFib = nMin1 + nMin2
			nMin2 = nMin1
			nMin1 = currentFib
		}
		return currentFib
	}
}
