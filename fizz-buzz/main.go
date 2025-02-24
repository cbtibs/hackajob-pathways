package main

import (
	"fmt"
	"strconv"
	"strings"
)

func run(N int, M int) string {
	if N > M {
		return "Error: N must be less than or equal to M"
	}

	var result strings.Builder

	for i := N; i <= M; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			result.WriteString("FizzBuzz")
		case i%3 == 0:
			result.WriteString("Fizz")
		case i%5 == 0:
			result.WriteString("Buzz")
		default:
			result.WriteString(strconv.Itoa(i))
		}

		if i < M {
			result.WriteString(",")
		}
	}

	return result.String()
}

func main() {
	fmt.Println(run(1, 5))
	fmt.Println(run(10, 15))
	fmt.Println(run(3, 3))
	fmt.Println(run(5, 5))
	fmt.Println(run(15, 15))
	fmt.Println(run(5, 1))
}
