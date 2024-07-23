package main

import (
	"fmt"
	"time"
)

func TimeTaken() {
	fmt.Println("---hello from timetaken---")
	t1 := time.Now()
	sum := 0
	for i := 0; i < 100_00_00_00_00; i++ {
		sum += i
	}
	fmt.Println(sum)
	t2 := time.Now()

	t := t2.Sub(t1)

	fmt.Println(t * 1000)
}
