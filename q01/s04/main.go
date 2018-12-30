package main

import (
	"fmt"
	"time"
)

const (
	totalNum       = 52
	maxCardsToPick = 5
	maxThreshold   = 22
	maxStart       = totalNum
)

var dp [maxCardsToPick + 1][maxThreshold + 1][maxStart + 1]int

// probability is equal to numOfPossiblecases divided by NumOfAllCases
func probability(cardsToPick, threshold int) float64 {
	var pokers [totalNum]int
	for i := 0; i < totalNum; i++ {
		pokers[i] = i/4 + 1
		if pokers[i] > 10 {
			pokers[i] = 10 // for J, Q, K
		}
	}
	for j := 0; j <= threshold; j++ {
		for k := 0; k <= totalNum; k++ {
			dp[0][j][k] = 1
		}
	}
	for i := 1; i <= cardsToPick; i++ {
		for j := 1; j <= threshold; j++ {
			for k := totalNum - 1; k >= 0; k-- {
				for l := k; l < totalNum; l++ {
					if j-pokers[l] < 0 {
						continue
					}
					dp[i][j][k] += dp[i-1][j-pokers[l]][l+1]
				}
			}
		}
	}
	possible := dp[cardsToPick][threshold][0]
	total := combination(totalNum, cardsToPick)
	return float64(possible) / float64(total)
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
