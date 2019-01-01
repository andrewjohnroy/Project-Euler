package main

import (
	"fmt"
	"math"
)

func main() {
	var nums [2000000]int
	for i := 0; i < 2000000; i++{
		nums[i] = i
	}
	nums[1] = 0
	var limit = int(math.Floor(math.Sqrt(2000000)))
	for i := 2; i <= limit; i++{
		var j = i * 2
		for j < 2000000{
			nums[j] = 0
			j += i
		}
	}
	var sum = 0
	for i := 0; i < 2000000; i++{
		sum += nums[i]
	}		
	fmt.Printf("%d", sum)
}
