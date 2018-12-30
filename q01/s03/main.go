package main

import (
	"fmt"
	"time"
)

const (
	totalNum = 52
)

// probability is equal to numOfPossiblecases divided by NumOfAllCases
func probability(cardsToPick, threshold int) float64 {
	var pokers [totalNum]int
	for i := 0; i < totalNum; i++ {
		pokers[i] = i/4 + 1
		if pokers[i] > 10 {
			pokers[i] = 10 // for J, Q, K
		}
	}
	status = make(map[statusKey]int)
	possible := numOfPossibleCases(cardsToPick, threshold, pokers, 0)
	total := combination(totalNum, cardsToPick)
	return float64(possible) / float64(total)
}

type statusKey struct {
	CardsToPick int
	Threshold   int
	Start       int
}

var status map[statusKey]int

// get number of possible cases via status cache
func numOfPossibleCases(cardsToPick, threshold int, pokers [totalNum]int, start int) (num int) {
	if threshold < 0 {
		return 0
	}
	if cardsToPick == 0 {
		return 1
	}
	if v, exist := status[statusKey{CardsToPick: cardsToPick, Threshold: threshold, Start: start}]; exist {
		return v
	}
	for i := start; i < len(pokers); i++ {
		num += numOfPossibleCases(cardsToPick-1, threshold-pokers[i], pokers, i+1)
	}
	status[statusKey{CardsToPick: cardsToPick, Threshold: threshold, Start: start}] = num
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
