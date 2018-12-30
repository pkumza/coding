package main

import (
	"fmt"
	"time"
)

const (
	totalNumOfPoker = 52
)

// probability is equal to numOfPossiblecases divided by NumOfAllCases
func probability(numOfCards, threshold int) float64 {
	var pokers [totalNumOfPoker]int
	for i := 0; i < totalNumOfPoker; i++ {
		pokers[i] = i/4 + 1
		if pokers[i] > 10 {
			pokers[i] = 10 // for J, Q, K
		}
	}
	possible := numOfPossibleCases(numOfCards, threshold, pokers[:])
	total := combination(totalNumOfPoker, numOfCards)
	return float64(possible) / float64(total)
}

// get number of possible cases
func numOfPossibleCases(numOfCards, threshold int, pokers []int) (num int) {
	if threshold < 0 {
		return 0
	}
	if numOfCards == 0 {
		return 1
	}
	for i := 0; i < len(pokers); i++ {
		num += numOfPossibleCases(numOfCards-1, threshold-pokers[i], pokers[i+1:])
	}
	return num
}

// C(x,y). this func have risk of overflow when totalNum is large
func combination(totalNum, subSetNum int) (result int) {
	result = 1
	for i := 0; i < subSetNum; i++ {
		result *= (totalNum - i)
	}
	for i := 1; i <= subSetNum; i++ {
		result /= i
	}
	return
}

func main() {
	jobStart := time.Now()
	fmt.Print("Result is \033[1;92m")
	fmt.Print(probability(5, 21))
	fmt.Println("\033[0m")
	fmt.Println("Time consumed", time.Since(jobStart))
}
