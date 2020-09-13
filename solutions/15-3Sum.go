package solutions

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	res := [][]int{}
	N := len(nums)

	// N
	if nums == nil || N < 3 {
		return res
	}

	counter := map[int]int{}

	for _, v := range nums {
		counter[v]++
	}

	unique := []int{}

	for k := range counter {
		unique = append(unique, k)
	}
	sort.Ints(unique)

	N = len(unique)

	for i := range unique {
		if unique[i] > 0 {
			return res
		}

		if unique[i] == 0 && counter[unique[i]] >= 3 {
			res = append(res, []int{0, 0, 0})
		}

		for j := i + 1; j < N; j++ {
			if unique[i]*2+unique[j] == 0 && counter[unique[i]] > 1 {
				res = append(res, []int{unique[i], unique[i], unique[j]})
			}

			if unique[j]*2+unique[i] == 0 && counter[unique[j]] > 1 {
				res = append(res, []int{unique[j], unique[j], unique[i]})
			}

			c := 0 - unique[i] - unique[j]

			if c > unique[i] && c > unique[j] && counter[c] > 0 {
				res = append(res, []int{unique[i], unique[j], c})
			}
		}
	}
	return res
}
