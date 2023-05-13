package future

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestAnonymousFunc(t *testing.T) {
	futureFunc := FutureFunc[int](Fibonacci, 10)
	result, err := futureFunc.Get()
	t.Logf("Result: %v\n", result)
	t.Logf("err: %v\n", err)
}

func TestFutureFunc(t *testing.T) {
	x := 10
	var elapsed time.Duration
	start := time.Now()
	future := FutureFunc[int](func() int {
		printTime(t)
		time.Sleep(5 * time.Second)
		fmt.Printf("x = %v\n", x)
		return x * 10
	})
	elapsed = time.Since(start)

	t.Logf("it took %s", elapsed)
	assert.Less(t, elapsed.Milliseconds(), (1 * time.Second).Milliseconds())

	result, _ := future.Get()
	elapsed = time.Since(start)
	assert.Less(t, (5 * time.Second).Milliseconds(), elapsed.Milliseconds())
	assert.Equal(t, 100, result)
	t.Logf("Result: %v\n", result)

	// This assert tests calling result second time doesn't cause any problems
	assert.Equal(t, 100, result)
}

func TestFutureFuncTimeOut(t *testing.T) {
	x := 10
	var elapsed time.Duration
	start := time.Now()
	future := FutureFunc[int](func() int {
		printTime(t)
		time.Sleep(5 * time.Second)
		fmt.Printf("x = %v\n", x)
		return x * 10
	})
	elapsed = time.Since(start)

	t.Logf("it took %s", elapsed)
	assert.Less(t, elapsed.Milliseconds(), (1 * time.Second).Milliseconds())

	result, err := future.GetWithTimeout(3 * time.Second)
	elapsed = time.Since(start)
	assert.Equal(t, nil, result)
	assert.Less(t, elapsed.Milliseconds(), (4 * time.Second).Milliseconds())
	t.Logf("Result: %v\n", result)
	t.Logf("err: %v\n", err.Error())
	assert.Equal(t, nil, result)
}

func printTime(t *testing.T) {
	t.Logf("time: %v\n", time.Now().Unix())
}
