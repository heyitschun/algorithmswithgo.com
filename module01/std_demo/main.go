package main

import (
	"algo/module01"
	"fmt"
)

func main() {
	var n int
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		var a, b int
		fmt.Scanf("%d %d", &a, &b)
		gcd := module01.GCD(a, b)
		fmt.Println(gcd)
	}
}

// If you are getting an error running: go run main.go < input.txt > output.txt
// and you use PowerShell, this is because the `<` operation is not supported
// in PowerShell (yet). Try using another command line interface.
