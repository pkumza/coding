package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	testTimes = 1000000
)

func probability(numOfCards, threshold int) float64 {
	pokers := make([]int, 52)
	for i := 0; i < 52; i++ {
		pokers[i] = min(i/4+1, 10)
	}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	match := 0
	for i := 0; i < testTimes; i++ {
		candidates := make(map[int]struct{})
		for len(candidates) < numOfCards {
			candidates[r.Intn(52)] = struct{}{}
		}
		sum := 0
		for j := range candidates {
			sum += pokers[j]
		}
		if sum <= threshold {
			match++
		}
	}
	return float64(match) / float64(testTimes)
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	jobStart := time.Now()
	fmt.Print("Result is \033[1;92m")
	fmt.Print(probability(5, 21))
	fmt.Println("\033[0m")
	fmt.Println("Time consumed", time.Since(jobStart))
}
