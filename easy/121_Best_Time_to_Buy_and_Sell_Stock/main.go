package main

import "fmt"

func main() {
	/*
		input cases
		{7,1,5,3,6,4}
		{7,6,4,3,1}
		{1,2}
	*/
	input := []int{7, 1, 5, 3, 6, 4}
	fmt.Println(maxProfit1(input))
}

// T: O(n^2), S: O(1)
func maxProfit(prices []int) int {
	max := 0

	for i, p := range prices {
		for j := i + 1; j < len(prices); j++ {
			tmp := prices[j] - p
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}

// T: O(n), S: O(1)
func maxProfit1(prices []int) int {
	if len(prices) == 0 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, p := range prices {
		if p < minPrice {
			minPrice = p
		} else {
			tmp := p - minPrice
			if tmp > maxProfit {
				maxProfit = tmp
			}
		}
	}
	return maxProfit
}
