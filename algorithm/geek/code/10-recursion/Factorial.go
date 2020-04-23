package recursion

import (
	"fmt"
)

// Fab 使用map记录阶乘的结果
type Fab struct {
	val map[int]int
}

// NewFab create a new Fab object
func NewFab(n int) *Fab {
	return &Fab{make(map[int]int, n)}
}

// Factorial compute the factorial of n
func (fab *Fab) Factorial(n int) int {
	if fab.val[n] != 0 {
		return fab.val[n]
	}
	if n <= 1 {
		fab.val[n] = 1
		return 1
	}
	res := n * fab.val[n-1]
	fab.val[n] = res
	return res
}

// Print print factorial of number n
func (fab *Fab) Print(n int) {
	fmt.Println(fab.val[n])
}
