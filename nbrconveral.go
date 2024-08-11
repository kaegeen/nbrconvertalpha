package main

import (
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]
	upper := false
	if len(args) > 0 && args[0] == "--upper" {
		upper = true
		args = args[1:]
	}

	toInt := func(s string) (int, bool) {
		result := 0
		for _, char := range s {
			if char < '0' || char > '9' {
				return 0, false
			}
			result = result*10 + int(char-'0')
		}
		return result, true
	}

	for _, arg := range args {

		n, valid := toInt(arg)
		if !valid || n < 1 || n > 26 {

			fmt.Print(" ")
		} else {

			var letter byte
			if upper {
				letter = 'A' + byte(n) - 1
			} else {
				letter = 'a' + byte(n) - 1
			}
			fmt.Print(string(letter))
		}
	}

	fmt.Println()
}
