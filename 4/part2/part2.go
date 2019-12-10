package main

import (
	"fmt"
	"strconv"
)

func hasSixDigits(number int) bool {
	numStr := strconv.Itoa(number)
	return len(numStr) == 6
}

func hasAscendingDigits(number int) bool {
	numStr := strconv.Itoa(number)
	for i := 0; i < (len(numStr) - 1); i++ {
		if numStr[i] > numStr[i+1] {
			return false
		}
	}
	return true
}

func digitHisto(number int) map[int]int {
	histo := make(map[int]int)
	numStr := strconv.Itoa(number)
	for _, rune := range numStr {
		digit, _ := strconv.Atoi(string(rune))
		histo[digit] = histo[digit] + 1
	}
	return histo
}

func hasExactlyTwoRepeatingDigits(number int) bool {
	histo := digitHisto(number)
	for _, count := range histo {
		if count == 2 {
			return true
		}
	}
	return false
}

func main() {
	const from = 100000
	const to = 200000

	count := 0

	for i := from; i <= to; i++ {
		if hasSixDigits(i) && hasAscendingDigits(i) && hasExactlyTwoRepeatingDigits(i) {
			count++
		}
	}

	fmt.Println("Number of possible passwords:", count)
}
