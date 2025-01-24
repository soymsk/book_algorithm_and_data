package main

import (
	"fmt"
	"math"
)

const INF int = math.MaxInt

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func main() {
	var S, T string
	fmt.Scan(&S)
	fmt.Scan(&T)

	dp := make([][]int, len(S)+1)
	for i := 0; i <= len(S); i++ {
		dp[i] = make([]int, len(T)+1)
		for j := 0; j <= len(T); j++ {
			dp[i][j] = INF
		}
		if i == 0 {
			dp[0][0] = 0
		}
		for j := 0; j <= len(T); j++ {
			if i > 0 && j > 0 {
				if S[i-1] == T[j-1] {
					// No ops
					dp[i][j] = min(dp[i][j], dp[i-1][j-1])
				} else {
					// Modify
					dp[i][j] = min(dp[i][j], dp[i-1][j-1]+1)
				}
			}

			// Deletion
			if i > 0 {
				dp[i][j] = min(dp[i][j], dp[i-1][j]+1)
			}

			// Addition
			if j > 0 {
				dp[i][j] = min(dp[i][j], dp[i][j-1]+1)
			}

		}
	}
	fmt.Println("Answer: ", dp[len(S)][len(T)])
}
