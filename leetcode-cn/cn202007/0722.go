package cn202007

import "fmt"

// Code0722 2020-07-22答案
func Code0722() {
	nums := []int{3, 4, 5, 6, 1, 2, 3}
	fmt.Println(fmt.Sprintf("Bisect result is %d", minArrayBisect(nums)))
	fmt.Println(fmt.Sprintf("Loop result is %d", minArrayLoop(nums)))
}

// minArrayBisect 二分法
func minArrayBisect(numbers []int) int {
	i, j := 0, len(numbers)-1
	mid := (i + j) / 2

	for i < j {
		if numbers[mid] > numbers[j] {
			i = mid + 1
		} else if numbers[mid] < numbers[j] {
			j = mid
		} else {
			j--
		}
		mid = (i + j) / 2
	}
	return numbers[i]
}

// minArrayLoop 循环法
func minArrayLoop(numbers []int) int {
	min := numbers[0]
	for _, n := range numbers {
		if n < min {
			return n
		}
	}

	return min
}
