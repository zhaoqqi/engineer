package recursion

import (
	"fmt"
)

// Fibs 使用map记录斐波那契数列的值
type Fibs struct {
	val map[int]int
}

// NewFibs create a new Fibs object
func NewFibs(n int) *Fibs {
	return &Fibs{make(map[int]int, n)}
}

// Fibonacci compute the fibonacci value of number n
func (fib *Fibs) Fibonacci(n int) int {
	if n <= 1 {
		fib.val[n] = 1
		return 1
	} else if n == 2 {
		fib.val[n] = 2
		return 2
	} else {
		res := fib.val[n-1] + fib.val[n-2]
		fib.val[n] = res
		return res
	}
}

//Print print fabonacci value of number n
func (fib *Fibs) Print(n int) {
	fmt.Println(fib.val[n])
}
