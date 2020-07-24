package cn202007

import "fmt"

// Code0724 除数博弈
func Code0724() {
	res := divisorGame(3)
	resStr := "true"
	if !res {
		resStr = "false"
	}

	fmt.Println(resStr)
}

func divisorGame(N int) bool {
	// 正解：return N % 2 == 0
	// 默认爱丽丝胜利
	ali, bob := true, true
	x := 0

	for N > 0 && ali && bob {
		x, ali = computeX(N)
		N = N - x
		if ali {
			x, bob = computeX(N)
			N = N - x
		}
	}

	return ali
}

// computeX 计算当前N的x
// retrn:x的值，是否成功
func computeX(N int) (int, bool) {
	if N == 1 {
		return 0, false
	}

	x := N - 1
	if N%x == 0 {
		return x, true
	}

	for ; x < N && x >= 1; x-- {
		if N%x == 0 && x%2 == 1 {
			return x, true
		}
	}

	return 0, false
}
