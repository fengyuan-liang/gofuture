// @Description: future func
// @Version: 1.0.0
// @Date: 2023/05/13 21:12
// @Author: fengyuan-liang@foxmail.com

package future

import "testing"

func TestSimpleFuture(t *testing.T) {
	future := SimpleFuture[int](func() (int, error) {
		return Fibonacci(10), nil
	})
	// blocking to wait result
	result, err := future()
	t.Logf("Result: %v\n", result)
	t.Logf("err: %v\n", err)
}

// Fibonacci returns the nth Fibonacci number
func Fibonacci(n int) int {
	if n <= 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}
