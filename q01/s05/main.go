package main

import (
	"fmt"
	"time"
)

const (
	totalNum       = 52
	maxCardsToPick = 5
	maxThreshold   = 21
	maxStart       = totalNum
)

var dp [maxCardsToPick][maxThreshold + 1][maxStart + 1]int

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
	for i := 1; i < cardsToPick; i++ {
		for j := 0; j <= threshold; j++ {
			for k := totalNum; k >= 0; k-- {
				dp[i%2][j][k] = 0
				for l := k; l < totalNum; l++ {
					if j-pokers[l] < 0 {
						continue
					}
					dp[i%2][j][k] += dp[(i-1)%2][j-pokers[l]][l+1]
				}
			}
		}
	}
	possible := 0
	for l := 0; l < totalNum; l++ {
		if threshold-pokers[l] >= 0 {
			possible += dp[(cardsToPick-1)%2][threshold-pokers[l]][l+1]
		}
	}
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
