package main

import (
	"fmt"
	"log"
)

func run(student_list []int) int {
	if len(student_list)%2 == 0 {
		log.Println("Warning: Input list length is even. There might be no unpaired student.")
	}
	single_student_number := 0
	for _, team := range student_list {
		single_student_number ^= team
	}
	return single_student_number
}

func main() {
	fmt.Println("Unpaired student:", run([]int{2, 4, 5, 4, 2}))
}
