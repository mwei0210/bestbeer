package main

import (
	"fmt"
)

const totalTap = 1000

func main() {
	for i := 1; i <= totalTap; i++ {
		var divs = []int{}

		//find all divisors
		for j := 1; j < totalTap; j++ {
			if i%j == 0 && i != j {
				divs = append(divs, j)
			}
		}
		//found all divisors

		//test condition 1 (must be more than tap number)
		if sum(divs) <= i {
			continue
		}

		//test condition 2 (no subset should add up to tap number)
		if !goodCombinations(divs, i) {
			continue
		}

		//print all the tap with best beer
		fmt.Println(i)
		fmt.Println(divs)

	}
}

func sum(divs []int) (total int) {
	for _, i := range divs {
		total += i
	}
	return
}

func goodCombinations(divs []int, tapNo int) bool {
	n := uint(len(divs))

	// computing all 2^n
	// subsets one by one
	for i := 0; i < (1 << n); i++ {
		comb := []int{}

		// compute current subset
		for j := uint(0); j < n; j++ {

			// (1<<j) is a number with jth bit 1
			// so when we bitwise 'and' them with the
			// subset number we get which numbers
			// are present in the subset and which
			// are not
			if (i & (1 << j)) > 0 {
				comb = append(comb, divs[j])
			}

		}
		if sum(comb) == tapNo {
			return false
		}
	}
	return true
}
