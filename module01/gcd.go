package module01

func GCD(a, b int) int {
	for {
		if a == 0 || b == 0 {
			if a > b {
				return a
			} else {
				return b
			}
		}
		a, b = b, a%b
	}
}
