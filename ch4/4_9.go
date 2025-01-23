package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// partial sum problem
func solve_partial_sum(n int, w int, a []int) bool {
	if n == 0 {
		if w == 0 {
			return true
		} else {
			return false
		}
	}

	if solve_partial_sum(n-1, w, a) {
		return true
	}

	if solve_partial_sum(n-1, w-a[n-1], a) {
		fmt.Println(a[n-1])
		return true

	}

	return false
}

func main() {

	sc := bufio.NewScanner(os.Stdin)

	sc.Scan()
	s := strings.Split(sc.Text(), " ")
	W, _ := strconv.Atoi(s[0])

	sc.Scan()
	strArr := strings.Split(sc.Text(), " ")

	N := len(strArr)
	a := make([]int, N)
	for i := range strArr {
		a[i], _ = strconv.Atoi(strArr[i])
	}

	fmt.Println(solve_partial_sum(N, W, a))

}
